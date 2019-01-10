package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert-client/pkg/client"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"net/http"
)

func HandlerResource(request *restful.Request, response *restful.Response) {
	conn, err := client.GetDispatcherGrpcLoadBalanceClient()

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	cli := pb.NewResourceHandlerClient(conn)

	method := request.Request.Method

	var rsp *pb.ResourceGroupResponse

	switch method {
	case http.MethodGet, http.MethodDelete:
		rgID := request.QueryParameter("resource_group_id")
		//recvID := request.QueryParameter("receiver_id")

		//if rgID == "" {
		//	response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ResourceGroupResponse{Error: &pb.Error{Text: "resource group id must be specified"}})
		//	return
		//}

		if method == http.MethodGet {
			rsp, _ = cli.GetResource(context.Background(), &pb.ResourceGroupSpec{ResourceGroupId: rgID})
		} else {
			rsp, _ = cli.DeleteResource(context.Background(), &pb.ResourceGroupSpec{ResourceGroupId: rgID})
		}

	case http.MethodPost, http.MethodPut:
		var resourceGroup = pb.ResourceGroup{}
		err := request.ReadEntity(&resourceGroup)

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ResourceGroupResponse{Error: &pb.Error{Text: err.Error()}})
			return
		}

		if method == http.MethodPost {
			rsp, _ = cli.CreateResource(context.Background(), &resourceGroup)
		} else {
			rsp, _ = cli.UpdateResource(context.Background(), &resourceGroup)
		}
	}

	if rsp != nil && rsp.Error.Code > 0 {
		glog.Errorln(rsp.Error)
		response.WriteHeaderAndEntity(http.StatusInternalServerError, rsp)
	} else {
		response.WriteHeaderAndEntity(http.StatusOK, rsp)
	}
}
