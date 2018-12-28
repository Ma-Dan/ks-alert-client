package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert-client/pkg/client"
	"github.com/carmanzhang/ks-alert-client/pkg/constant"
	"github.com/carmanzhang/ks-alert/pkg/dispatcher/pb"
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"net/http"
)

func CreateResourceType(request *restful.Request, response *restful.Response) {

}
func RetrieveResourceType(request *restful.Request, response *restful.Response) {

}
func UpdateResourceType(request *restful.Request, response *restful.Response) {

}
func DeleteResourceType(request *restful.Request, response *restful.Response) {

}

func HandlerResourceType(request *restful.Request, response *restful.Response) {
	conn, err := client.GetDispatcherGrpcLoadBalanceClient()

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	cli := pb.NewResourceTypeHandlerClient(conn)

	method := request.Request.Method

	switch method {
	case http.MethodGet, http.MethodDelete:
		typeID := request.QueryParameter("resource_type_id")
		// qingcloud + kubesphere + resource_type_name
		typeName := request.QueryParameter("resource_type_name")

		var resourceTypeSpec = pb.ResourceTypeSpec{
			ResourceTypeId:   typeID,
			ResourceTypeName: typeName,
			ProductId:        GetProductID(conn),
		}

		if method == http.MethodGet {
			rsp, err := cli.GetResourceType(context.Background(), &resourceTypeSpec)

			if err != nil {
				glog.Errorln(err)
				response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ResourceTypeResponse{Error: &pb.Error{Text: err.Error()}})
				return
			} else {
				response.WriteHeaderAndEntity(http.StatusOK, rsp)
			}
		} else {
			rsp, err := cli.DeleteResourceType(context.Background(), &resourceTypeSpec)

			if err != nil {
				glog.Errorln(err)
				response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ResourceTypeResponse{Error: &pb.Error{Text: err.Error()}})
				return
			} else {
				response.WriteHeaderAndEntity(http.StatusOK, rsp)
			}
		}

	case http.MethodPost, http.MethodPut:
		var resourceType = pb.ResourceType{}
		err := request.ReadEntity(&resourceType)

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ResourceTypeResponse{Error: &pb.Error{Text: err.Error()}})
			return
		}

		if resourceType.ProductId == "" {
			resourceType.ProductId = GetProductID(conn)
		}

		if method == http.MethodPost {
			rsp, err := cli.CreateResourceType(context.Background(), &resourceType)

			if err != nil {
				glog.Errorln(err)
				response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ResourceTypeResponse{Error: &pb.Error{Text: err.Error()}})
				return
			} else {
				response.WriteHeaderAndEntity(http.StatusOK, rsp)
			}
		} else {

			rsp, err := cli.UpdateResourceType(context.Background(), &resourceType)

			if err != nil {
				glog.Errorln(err)
				response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ResourceTypeResponse{Error: &pb.Error{Text: err.Error()}})
				return
			} else {
				response.WriteHeaderAndEntity(http.StatusOK, rsp)
			}
		}
	}
}

// GetProductID by resource_type_name + enterprise_name + product_name
func GetProductID(conn *grpc.ClientConn) string {
	cli := pb.NewProductHandlerClient(conn)
	prod, err := cli.GetProduct(context.Background(), &pb.ProductSpec{
		ProductName: constant.Product,
	})
	if err != nil {
		glog.Error(err.Error())
		return ""
	}
	return prod.Product.ProductId
}
