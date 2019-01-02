package handler

import (
	"context"
	"fmt"
	"github.com/carmanzhang/ks-alert-client/pkg/client"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"github.com/carmanzhang/ks-alert/pkg/utils/jsonutil"
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
		//recvID := request.QueryParameter("receiver_id")

		if recvGroupID == "" {
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ReceiverGroupResponse{Error: &pb.Error{Text: "receiver group id must be specified"}})
			return
		}

		var rsp *pb.ReceiverGroupResponse

		if method == http.MethodGet {
			rsp, err = cli.GetReceiver(context.Background(), &pb.ReceiverGroupSpec{ReceiverGroupId: recvGroupID})
		} else {
			rsp, err = cli.DeleteReceiver(context.Background(), &pb.ReceiverGroupSpec{ReceiverGroupId: recvGroupID})
		}

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ReceiverGroupResponse{Error: &pb.Error{Text: err.Error()}})
			return
		} else {
			fmt.Println(jsonutil.Marshal(rsp))
			response.WriteHeaderAndEntity(http.StatusOK, rsp)
			return
		}

	case http.MethodPost, http.MethodPut:
		var receiverGroup = pb.ReceiverGroup{}
		err := request.ReadEntity(&receiverGroup)

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ReceiverGroupResponse{Error: &pb.Error{Text: err.Error()}})
			return
		}

		var rsp *pb.ReceiverGroupResponse

		if method == http.MethodPost {
			rsp, err = cli.CreateReceiver(context.Background(), &receiverGroup)
		} else {
			rsp, err = cli.UpdateReceiver(context.Background(), &receiverGroup)
		}

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ReceiverGroupResponse{Error: &pb.Error{Text: err.Error()}})
			return
		} else {
			response.WriteHeaderAndEntity(http.StatusOK, rsp)
		}
	}
}
