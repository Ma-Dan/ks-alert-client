package handler

import (
	"context"
	"kubesphere.io/ks-alert-client/pkg/client"
	"kubesphere.io/ks-alert/pkg/pb"
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

	var rsp *pb.ReceiverGroupResponse

	switch method {
	case http.MethodGet, http.MethodDelete:

		recvGroupID := request.QueryParameter("receiver_group_id")
		//recvID := request.QueryParameter("receiver_id")

		if method == http.MethodGet {
			rsp, _ = cli.GetReceiver(context.Background(), &pb.ReceiverGroupSpec{ReceiverGroupId: recvGroupID})
		} else {
			rsp, _ = cli.DeleteReceiver(context.Background(), &pb.ReceiverGroupSpec{ReceiverGroupId: recvGroupID})
		}

	case http.MethodPost, http.MethodPut:
		var receiverGroup = pb.ReceiverGroup{}
		err := request.ReadEntity(&receiverGroup)

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ReceiverGroupResponse{Error: &pb.Error{Text: err.Error()}})
			return
		}

		if method == http.MethodPost {
			rsp, _ = cli.CreateReceiver(context.Background(), &receiverGroup)
		} else {
			rsp, _ = cli.UpdateReceiver(context.Background(), &receiverGroup)
		}
	}

	if rsp != nil && rsp.Error.Code > 0 {
		glog.Errorln(rsp.Error)
		response.WriteHeaderAndEntity(http.StatusInternalServerError, rsp)
	} else {
		response.WriteHeaderAndEntity(http.StatusOK, rsp)
	}
}
