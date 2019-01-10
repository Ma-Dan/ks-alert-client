package handler

import (
	"context"
	"github.com/carmanzhang/ks-alert-client/pkg/client"
	"github.com/carmanzhang/ks-alert-client/pkg/constant"
	"github.com/carmanzhang/ks-alert/pkg/pb"
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"net/http"
)

//func CreateProduct(request *restful.Request, response *restful.Response) {
//	var prod = pb.Product{}
//	err := request.ReadEntity(&prod)
//
//	if err != nil {
//		glog.Errorln(err)
//		response.WriteError(http.StatusInternalServerError,err)
//		return
//	}
//
//	conn, err := client.GetDispatcherGrpcLoadBalanceClient()
//
//	if err != nil {
//		glog.Errorln(err)
//		response.WriteError(http.StatusInternalServerError,err)
//		return
//	}
//
//	cli := pb.NewProductHandlerClient(conn)
//	rsp, err := cli.CreateProduct(context.Background(), &prod)
//
//	if err != nil {
//		glog.Errorln(err)
//		response.WriteError(http.StatusInternalServerError,err)
//		return
//	}else {
//		response.WriteHeaderAndEntity(http.StatusOK, rsp)
//	}
//}
//
//func RetrieveProduct(request *restful.Request, response *restful.Response) {
//	prodID := request.QueryParameter("product_id")
//	prodName := request.QueryParameter("product_name")
//
//	conn, err := client.GetDispatcherGrpcLoadBalanceClient()
//
//	if err != nil {
//		glog.Errorln(err)
//		response.WriteError(http.StatusInternalServerError,err)
//		return
//	}
//
//	cli := pb.NewProductHandlerClient(conn)
//
//	rsp, err := cli.GetProduct(context.Background(), &pb.ProductSpec{
//		ProductId:   prodID,
//		ProductName: prodName,
//	})
//
//	if err != nil {
//		glog.Errorln(err)
//		response.WriteError(http.StatusInternalServerError,err)
//		return
//	}else {
//		response.WriteHeaderAndEntity(http.StatusOK, rsp)
//	}
//}
//
//func UpdateProduct(request *restful.Request, response *restful.Response) {
//	var prod = pb.Product{}
//	err := request.ReadEntity(&prod)
//	if err != nil {
//		glog.Errorln(err)
//		response.WriteError(http.StatusInternalServerError, err)
//		return
//	}
//
//	if prod.ProductId == "" && prod.ProductName == "" {
//		response.WriteHeaderAndEntity(http.StatusOK, errors.New("product_id or product_name must be specified"))
//		return
//	}
//
//	conn, err := client.GetDispatcherGrpcLoadBalanceClient()
//
//	if err != nil {
//		glog.Errorln(err)
//		response.WriteError(http.StatusInternalServerError,err)
//		return
//	}
//
//	cli := pb.NewProductHandlerClient(conn)
//
//	rsp, err := cli.UpdateProduct(context.Background(), &prod)
//
//	if err != nil {
//		glog.Errorln(err)
//		response.WriteError(http.StatusInternalServerError,err)
//		return
//	}else {
//		response.WriteHeaderAndEntity(http.StatusOK, rsp)
//	}
//}
//
//func DeleteProduct(request *restful.Request, response *restful.Response) {
//	prodID := request.QueryParameter("product_id")
//	prodName := request.QueryParameter("product_name")
//	conn, err := client.GetDispatcherGrpcLoadBalanceClient()
//
//	if err != nil {
//		glog.Errorln(err)
//		response.WriteError(http.StatusInternalServerError,err)
//		return
//	}
//
//	cli := pb.NewProductHandlerClient(conn)
//
//	rsp, err := cli.DeleteProduct(context.Background(), &pb.ProductSpec{
//		ProductId:   prodID,
//		ProductName: prodName,
//	})
//
//	if err != nil {
//		glog.Errorln(err)
//		response.WriteError(http.StatusInternalServerError,err)
//		return
//	}else {
//		response.WriteHeaderAndEntity(http.StatusOK, rsp)
//	}
//}

func HandlerProduct(request *restful.Request, response *restful.Response) {
	conn, err := client.GetDispatcherGrpcLoadBalanceClient()

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	cli := pb.NewProductHandlerClient(conn)

	method := request.Request.Method

	switch method {
	case http.MethodGet, http.MethodDelete:
		prodID := request.QueryParameter("product_id")
		prodName := request.QueryParameter("product_name")

		if method == http.MethodGet {
			rsp, err := cli.GetProduct(context.Background(), &pb.ProductSpec{
				ProductId:   prodID,
				ProductName: prodName,
			})

			if err != nil {
				glog.Errorln(err)
				response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ProductResponse{Error: &pb.Error{Text: err.Error()}})
				return
			} else {
				response.WriteHeaderAndEntity(http.StatusOK, rsp)
			}
		} else {
			rsp, err := cli.DeleteProduct(context.Background(), &pb.ProductSpec{
				ProductId:   prodID,
				ProductName: prodName,
			})

			if err != nil {
				glog.Errorln(err)
				response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ProductResponse{Error: &pb.Error{Text: err.Error()}})
				return
			} else {
				response.WriteHeaderAndEntity(http.StatusOK, rsp)
			}
		}
	case http.MethodPost, http.MethodPut:
		var prod = pb.Product{}
		err := request.ReadEntity(&prod)

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ProductResponse{Error: &pb.Error{Text: err.Error()}})
			return
		}

		if prod.EnterpriseName == "" && prod.EnterpriseId == "" {
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ProductResponse{Error: &pb.Error{Text: "enterprise_id or enterprise_name must be specified"}})
			return
		}

		if prod.EnterpriseId == "" {
			cliEnt := pb.NewEnterpriseHandlerClient(conn)
			ent, err := cliEnt.GetEnterprise(context.Background(), &pb.EnterpriseSpec{EnterpriseId: prod.EnterpriseName})
			if err != nil {
				glog.Error(err.Error())
				response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ProductResponse{Error: &pb.Error{Text: err.Error()}})
				return
			}
			prod.EnterpriseId = ent.Enterprise.EnterpriseId
		}

		if method == http.MethodPost {
			rsp, err := cli.CreateProduct(context.Background(), &prod)

			if err != nil {
				glog.Errorln(err)
				response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ProductResponse{Error: &pb.Error{Text: err.Error()}})
				return
			} else {
				response.WriteHeaderAndEntity(http.StatusOK, rsp)
			}
		} else {

			if prod.ProductId == "" && prod.ProductName == "" {
				response.WriteHeaderAndEntity(http.StatusOK, &pb.ProductResponse{Error: &pb.Error{Text: "product_id or product_name must be specified"}})
				return
			}
			rsp, err := cli.UpdateProduct(context.Background(), &prod)

			if err != nil {
				glog.Errorln(err)
				response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.ProductResponse{Error: &pb.Error{Text: err.Error()}})
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
