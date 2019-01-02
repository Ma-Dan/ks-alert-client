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

func CreateResource(request *restful.Request, response *restful.Response) {

}

func RetrieveResource(request *restful.Request, response *restful.Response) {

}

func UpdateResource(request *restful.Request, response *restful.Response) {

}

func DeleteResource(request *restful.Request, response *restful.Response) {

}

func HandlerResource(request *restful.Request, response *restful.Response) {
	conn, err := client.GetDispatcherGrpcLoadBalanceClient()

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	cli := pb.NewResourceHandlerClient(conn)

	method := request.Request.Method

	switch method {
	case http.MethodGet, http.MethodDelete:
		rgID := request.QueryParameter("resource_group_id")
		//recvID := request.QueryParameter("receiver_id")

		if rgID == "" {
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ResourceGroupResponse{Error: &pb.Error{Text: "resource group id must be specified"}})
			return
		}

		var rsp *pb.ResourceGroupResponse

		if method == http.MethodGet {
			rsp, err = cli.GetResource(context.Background(), &pb.ResourceGroupSpec{ResourceGroupId: rgID})
		} else {
			rsp, err = cli.DeleteResource(context.Background(), &pb.ResourceGroupSpec{ResourceGroupId: rgID})
		}

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ResourceGroupResponse{Error: &pb.Error{Text: err.Error()}})
			return
		} else {
			fmt.Println(jsonutil.Marshal(rsp))
			response.WriteHeaderAndEntity(http.StatusOK, rsp)
			return
		}

	case http.MethodPost, http.MethodPut:
		var resourceGroup = pb.ResourceGroup{}
		err := request.ReadEntity(&resourceGroup)

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ResourceGroupResponse{Error: &pb.Error{Text: err.Error()}})
			return
		}

		var rsp *pb.ResourceGroupResponse

		if method == http.MethodPost {
			rsp, err = cli.CreateResource(context.Background(), &resourceGroup)
		} else {
			rsp, err = cli.UpdateResource(context.Background(), &resourceGroup)
		}

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ResourceGroupResponse{Error: &pb.Error{Text: err.Error()}})
			return
		} else {
			response.WriteHeaderAndEntity(http.StatusOK, rsp)
		}
	}
}
