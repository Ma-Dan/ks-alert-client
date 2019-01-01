package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert-client/pkg/client"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"net/http"
)

func HandlerReceiver(request *restful.Request, response *restful.Response) {
	conn, err := client.GetDispatcherGrpcLoadBalanceClient()

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	cli := pb.NewReceiverHandlerClient(conn)

	method := request.Request.Method

	switch method {
	case http.MethodGet, http.MethodDelete:

		recvGroupID := request.QueryParameter("receiver_group_id")

		if method == http.MethodGet {
			rsp, err := cli.GetReceiver(context.Background(), &pb.ReceiverGroupSpec{ReceiverGroupId: recvGroupID})
			if err != nil {
				glog.Errorln(err)
				response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ResourceTypeResponse{Error: &pb.Error{Text: err.Error()}})
				return
			} else {
				response.WriteHeaderAndEntity(http.StatusOK, rsp)
			}

		} else {

		}

	case http.MethodPost, http.MethodPut:
		var receiverGroup = pb.ReceiverGroup{}
		err := request.ReadEntity(&receiverGroup)

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ResourceTypeResponse{Error: &pb.Error{Text: err.Error()}})
			return
		}

		if method == http.MethodPost {
			rsp, err := cli.CreateReceiver(context.Background(), &receiverGroup)

			if err != nil {
				glog.Errorln(err)
				response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ResourceTypeResponse{Error: &pb.Error{Text: err.Error()}})
				return
			} else {
				response.WriteHeaderAndEntity(http.StatusOK, rsp)
			}
		} else {
		}
	}
}
