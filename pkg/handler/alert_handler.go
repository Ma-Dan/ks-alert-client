package handler

import (
	"context"
	"fmt"
	. "github.com/carmanzhang/ks-alert-client/pkg/client"
	"github.com/carmanzhang/ks-alert-client/pkg/models"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
)

func CreateAlert(request *restful.Request, response *restful.Response) {

	var alertConfig pb.AlertConfig
	err := request.ReadEntity(&alertConfig)
	if err != nil {
		glog.Errorln(err)
	}

	conn, err := GetDispatcherGrpcLoadBalanceClient()
	if err != nil {
		fmt.Println(err)
	}

	//defer conn.Close()

	client := pb.NewAlertConfigHandlerClient(conn)
	//
	alertConfigResponse, err := client.CreateAlertConfig(context.Background(), &alertConfig)
	fmt.Println(alertConfigResponse, err)

}

func RetrieveAlert(request *restful.Request, response *restful.Response) {
	// find alert_configs by user_id
	userID := request.QueryParameter("user_id")
	// find alert_config by alert_config_id
	alertConfigID := request.QueryParameter("alert_config_id")

	conn, err := GetDispatcherGrpcLoadBalanceClient()
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	client := pb.NewAlertConfigHandlerClient(conn)

	var alertConfigIDs []string
	if alertConfigID != "" {
		alertConfigIDs = append(alertConfigIDs, alertConfigID)
	} else if userID != "" {
		userAlerts, err := models.GetAlertByUserID(userID)
		if err != nil {
			glog.Errorln(err.Error())
		}
		alertConfigIDs = models.GetAlertConfigIDs(userAlerts)

	} else {
		// find alert_configs by resource_type and  resource name
		resourceType := models.ResourceType(request.QueryParameter("resource_type"))
		resourceName := request.QueryParameter("resource_name")
		cluster := request.QueryParameter("cluster")
		node := request.QueryParameter("node")
		workspace := request.QueryParameter("workspace")
		namespace := request.QueryParameter("namespace")
		workload := request.QueryParameter("workload")
		//pod := request.QueryParameter("pod")

		var userAlert = models.UserAlertBinding{}
		userAlert.ResourceName = resourceName

		switch resourceType {
		case models.Cluster:
			if cluster != "" {
				userAlert.Cluster = "local"
			}
		case models.Node:
			if node != "" {
				userAlert.Node = node
			}
		case models.Workspace:
			if workspace != "" {
				userAlert.Workspace = workspace
			}
		case models.Namespace:
			if namespace != "" {
				userAlert.Workspace = workspace
				userAlert.Namespace = namespace
			}
		case models.Workload:
			if namespace != "" && workload != "" {
				userAlert.Namespace = namespace
				userAlert.Workload = workload
			}
		}
		userAlerts, err := models.GetAlertByResourceName(&userAlert)
		if err != nil {
			glog.Errorln(err.Error())
		}
		alertConfigIDs = models.GetAlertConfigIDs(userAlerts)
	}

	for _, acID := range alertConfigIDs {
		alertConfigResponse, err := client.GetAlertConfig(context.Background(), &pb.AlertConfig{AlertConfigId: acID})
		if err != nil {
			glog.Error(err.Error())
		}
		if alertConfigResponse.Error != nil {
			glog.Error(alertConfigResponse.Error.Text)
		}

		alertConfigResponse.GetAlertConfigId()
		alertConfigResponse.GetAlertConfig()
		alertConfigResponse.GetError()

	}
}

func UpdateAlert(request *restful.Request, response *restful.Response) {

}

func DeleteAlert(request *restful.Request, response *restful.Response) {

}

func RetrieveAlertHistory(request *restful.Request, response *restful.Response) {

}
func ListAlertRules(request *restful.Request, response *restful.Response) {

}
