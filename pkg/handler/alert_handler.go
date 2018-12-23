package handler

import (
	"context"
	"fmt"
	"github.com/emicklei/go-restful"
	"kubesphere.io/ks-alert/pkg/dispatcher/client"
	"kubesphere.io/ks-alert/pkg/dispatcher/option"
	executor "kubesphere.io/ks-alert/pkg/executor/pb"
	"kubesphere.io/ks-alert/pkg/models"
)

func CreateAlert(request *restful.Request, response *restful.Response) {

	//var alertConfig models.AlertConfig
	//err := request.ReadEntity(&alertConfig)
	//if err != nil {
	//	glog.Errorln(err)
	//}
	//
	//alertRuleGroup, err := models.CreateAlertRuleGroup(&alertConfig.AlertRuleGroup)
	//
	//// alertRules
	//_, err = models.CreateAlertRules(&alertConfig.AlertRuleGroup.AlertRules, alertRuleGroup.AlertRuleGroupID)
	//if err != nil {
	//	glog.Errorln(err)
	//}
	//
	//resourceGroup, err := models.CreateResourceGroup(alertConfig.ResourceGroup.ResourceGroupName, alertConfig.ResourceGroup.Description)
	//if err != nil {
	//	glog.Errorln(err)
	//}
	//
	//err = models.CreateResources(&alertConfig.ResourceGroup.Resources, resourceGroup, &alertConfig.URIParams)
	//if err != nil {
	//	glog.Errorln(err)
	//}
	//
	//receiverGroup, err := models.CreateReceiverGroup(&alertConfig.ReceiverGroup)
	//if err != nil {
	//	glog.Errorln(err)
	//}
	//
	//receivers, err := models.CreateReceivers(&alertConfig.ReceiverGroup.Receivers)
	//if err != nil {
	//	glog.Errorln(err)
	//}
	//
	//err = models.CreateReceiverBindingGroupItem(receivers, receiverGroup)
	//if err != nil {
	//	glog.Errorln(err)
	//}

	clientConn, err := client.GetExecutorGrpcLoadBalancerClient(*option.ExecutorServiceName, *option.EtcdAddr)

	if err != nil {
		panic(err)
	}

	// get banding host by
	client11 := executor.NewExecutorClient(clientConn)

	resp, err := client11.ExecuteAlertConfig(context.Background(), &executor.AlertConfig{Signal: executor.AlertConfig_Signal(models.Create), AlertConfigId: "world"})
	fmt.Println(resp)
	if err != nil {
		fmt.Println(err)
	}

}

func RetrieveAlert(request *restful.Request, response *restful.Response) {

}

func UpdateAlert(request *restful.Request, response *restful.Response) {

}

func DeleteAlert(request *restful.Request, response *restful.Response) {

}

func ListAlertRules(request *restful.Request, response *restful.Response) {

}
