package handler

import (
	"context"
	"fmt"
	"github.com/carmanzhang/ks-alert-client/pkg/client"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"net/http"
)

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
			rsp, err = cli.CreateAlertConfig(context.Background(), &alertConfig)
		} else {
			rsp, err = cli.UpdateAlertConfig(context.Background(), &alertConfig)

		}
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	if rsp != nil && rsp.Error.Code > 0 {
		glog.Errorln(rsp.Error)
		response.WriteHeaderAndEntity(http.StatusInternalServerError, rsp)
	} else {
		response.WriteHeaderAndEntity(http.StatusOK, rsp)
	}
}
