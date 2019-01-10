package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert-client/pkg/client"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"net/http"
)

func HandlerSuggestion(request *restful.Request, response *restful.Response) {
	conn, err := client.GetDispatcherGrpcLoadBalanceClient()

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	cli := pb.NewSuggestionHandlerClient(conn)

	method := request.Request.Method

	switch method {
	case http.MethodGet:

		configID := request.QueryParameter("alert_config_id")
		ruleID := request.QueryParameter("alert_rule_id")
		resourceID := request.QueryParameter("resource_id")

		rsp, err := cli.GetSuggestion(context.Background(), &pb.Suggestion{AlertConfigId: configID, AlertRuleId: ruleID, ResourceId: resourceID})

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.SuggestionResponse{Error: &pb.Error{Text: err.Error()}})
			return
		} else {
			response.WriteHeaderAndEntity(http.StatusOK, rsp)
			return
		}

	case http.MethodPut:
		var sugg = pb.Suggestion{}
		err := request.ReadEntity(&sugg)

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.SuggestionResponse{Error: &pb.Error{Text: err.Error()}})
			return
		}

		rsp, err := cli.UpdateSuggestion(context.Background(), &sugg)

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.SuggestionResponse{Error: &pb.Error{Text: err.Error()}})
			return
		} else {
			response.WriteHeaderAndEntity(http.StatusOK, rsp)
		}
	}
}
