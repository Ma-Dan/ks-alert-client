package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert-client/pkg/client"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"net/http"
)

func HandlerSilence(request *restful.Request, response *restful.Response) {
	conn, err := client.GetDispatcherGrpcLoadBalanceClient()

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	cli := pb.NewSilenceHandlerClient(conn)

	method := request.Request.Method

	var rsp *pb.SilenceResponse

	switch method {
	case http.MethodGet, http.MethodDelete:
		ruleID := request.QueryParameter("alert_rule_id")
		resID := request.QueryParameter("resource_id")

		var silence = pb.Silence{
			AlertRuleId: ruleID,
			ResourceId:  resID,
		}

		if method == http.MethodGet {
			rsp, err = cli.GetSilence(context.Background(), &silence)
		} else {
			rsp, err = cli.DeleteSilence(context.Background(), &silence)
		}

	case http.MethodPost, http.MethodPut:
		var silence = pb.Silence{}
		err := request.ReadEntity(&silence)

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.SilenceResponse{Error: &pb.Error{Text: err.Error()}})
			return
		}

		if silence.AlertRuleId == "" || silence.ResourceId == "" {
			errStr := "rule id and resource id must be specified"
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.SilenceResponse{Error: &pb.Error{Text: errStr}})
			return
		}

		if method == http.MethodPost {
			rsp, err = cli.CreateSilence(context.Background(), &silence)

		} else {
			rsp, err = cli.UpdateSilence(context.Background(), &silence)
		}
	}

	if err != nil {
		glog.Errorln(err)
		response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.SilenceResponse{Error: &pb.Error{Text: err.Error()}})
		return
	} else {
		response.WriteHeaderAndEntity(http.StatusOK, rsp)
	}

}
