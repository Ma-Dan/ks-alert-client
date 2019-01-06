package handler

import (
	"context"
	"fmt"
	"github.com/carmanzhang/ks-alert-client/pkg/client"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"net/http"
)

func RetrieveAlert(request *restful.Request, response *restful.Response) {
	//// find alert_configs by user_id
	//userID := request.QueryParameter("user_id")
	//// find alert_config by alert_config_id
	//alertConfigID := request.QueryParameter("alert_config_id")
	//
	//conn, err := GetDispatcherGrpcLoadBalanceClient()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//defer conn.Close()
	//
	//client := pb.NewAlertConfigHandlerClient(conn)
	//
	//var alertConfigIDs []string
	//if alertConfigID != "" {
	//	alertConfigIDs = append(alertConfigIDs, alertConfigID)
	//} else if userID != "" {
	//	userAlerts, err := models.GetAlertByUserID(userID)
	//	if err != nil {
	//		glog.Errorln(err.Error())
	//	}
	//	alertConfigIDs = models.GetAlertConfigIDs(userAlerts)
	//
	//} else {
	//	// find alert_configs by resource_type and  resource name
	//	resourceType := models.ResourceType(request.QueryParameter("resource_type"))
	//	resourceName := request.QueryParameter("resource_name")
	//	cluster := request.QueryParameter("cluster")
	//	node := request.QueryParameter("node")
	//	workspace := request.QueryParameter("workspace")
	//	namespace := request.QueryParameter("namespace")
	//	workload := request.QueryParameter("workload")
	//	//pod := request.QueryParameter("pod")
	//
	//	var userAlert = models.UserAlertBinding{}
	//	userAlert.ResourceName = resourceName
	//
	//	switch resourceType {
	//	case models.Cluster:
	//		if cluster != "" {
	//			userAlert.Cluster = "local"
	//		}
	//	case models.Node:
	//		if node != "" {
	//			userAlert.Node = node
	//		}
	//	case models.Workspace:
	//		if workspace != "" {
	//			userAlert.Workspace = workspace
	//		}
	//	case models.Namespace:
	//		if namespace != "" {
	//			userAlert.Workspace = workspace
	//			userAlert.Namespace = namespace
	//		}
	//	case models.Workload:
	//		if namespace != "" && workload != "" {
	//			userAlert.Namespace = namespace
	//			userAlert.Workload = workload
	//		}
	//	}
	//	userAlerts, err := models.GetAlertByResourceName(&userAlert)
	//	if err != nil {
	//		glog.Errorln(err.Error())
	//	}
	//	alertConfigIDs = models.GetAlertConfigIDs(userAlerts)
	//}
	//
	//for _, acID := range alertConfigIDs {
	//	alertConfigResponse, err := client.GetAlertConfig(context.Background(), &pb.AlertConfig{AlertConfigId: acID})
	//	if err != nil {
	//		glog.Error(err.Error())
	//	}
	//	if alertConfigResponse.Error != nil {
	//		glog.Error(alertConfigResponse.Error.Text)
	//	}
	//
	//	alertConfigResponse.GetAlertConfigId()
	//	alertConfigResponse.GetAlertConfig()
	//	alertConfigResponse.GetError()
	//
	//}
}

func RetrieveAlertHistory(request *restful.Request, response *restful.Response) {

}

func HandlerAlertConfig(request *restful.Request, response *restful.Response) {
	conn, err := client.GetDispatcherGrpcLoadBalanceClient()

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	cli := pb.NewAlertConfigHandlerClient(conn)

	method := request.Request.Method

	var rsp *pb.AlertConfigResponse

	switch method {
	case http.MethodGet, http.MethodDelete:
		configID := request.QueryParameter("alert_config_id")

		if method == http.MethodGet {
			rsp, _ = cli.GetAlertConfig(context.Background(), &pb.AlertConfigSpec{AlertConfigId: configID})
		} else {
			rsp, err = cli.DeleteAlertConfig(context.Background(), &pb.AlertConfigSpec{AlertConfigId: configID})
			if err != nil {
				fmt.Print(err)
			}
		}

	case http.MethodPost, http.MethodPut:
		var alertConfig = pb.AlertConfig{}
		err := request.ReadEntity(&alertConfig)

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.AlertConfigResponse{Error: &pb.Error{Text: err.Error()}})
			return
		}

		if method == http.MethodPost {
			rsp, _ = cli.CreateAlertConfig(context.Background(), &alertConfig)
		} else {
			rsp, _ = cli.UpdateAlertConfig(context.Background(), &alertConfig)

		}
	}

	if rsp != nil && rsp.Error.Code > 0 {
		glog.Errorln(rsp.Error)
		response.WriteHeaderAndEntity(http.StatusInternalServerError, rsp)
	} else {
		response.WriteHeaderAndEntity(http.StatusOK, rsp)
	}
}
