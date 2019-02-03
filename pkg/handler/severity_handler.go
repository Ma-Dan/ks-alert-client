package handler

import (
	"context"
	"kubesphere.io/ks-alert-client/pkg/client"
	"kubesphere.io/ks-alert-client/pkg/constant"
	"kubesphere.io/ks-alert/pkg/pb"
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"net/http"
)

func HandlerSeverity(request *restful.Request, response *restful.Response) {
	conn, err := client.GetDispatcherGrpcLoadBalanceClient()

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	cli := pb.NewSeverityHandlerClient(conn)

	method := request.Request.Method

	var rsp *pb.SeverityResponse

	switch method {
	case http.MethodGet, http.MethodDelete:
		severityID := request.QueryParameter("severity_id")
		// list severities of specific product
		prodID := request.QueryParameter("product_id")

		var sevSpec = pb.SeveritySpec{
			SeverityId:  severityID,
			ProductName: constant.Product,
			ProductId:   prodID,
		}

		if method == http.MethodGet {
			rsps, err := cli.GetSeverity(context.Background(), &sevSpec)
			if err != nil {
				glog.Errorln(err)
				response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.SeveritiesResponse{Error: &pb.Error{Text: err.Error()}})
				return
			} else {
				response.WriteHeaderAndEntity(http.StatusOK, rsps)
				return
			}

		} else {
			rsp, err = cli.DeleteSeverity(context.Background(), &sevSpec)
		}

	case http.MethodPost, http.MethodPut:
		var severity = pb.Severity{}
		err := request.ReadEntity(&severity)

		if err != nil {
			glog.Errorln(err)
			response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.SeverityResponse{Error: &pb.Error{Text: err.Error()}})
			return
		}

		if method == http.MethodPost {

			if severity.SeverityCh == "" || severity.SeverityEn == "" {
				errStr := "severity name must be specified"
				response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.SeverityResponse{Error: &pb.Error{Text: errStr}})
				return
			}

			rsp, err = cli.CreateSeverity(context.Background(), &severity)

		} else {
			if severity.SeverityId == "" {
				errStr := "severity id must be specified"
				response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.SeverityResponse{Error: &pb.Error{Text: errStr}})
				return
			}

			rsp, err = cli.UpdateSeverity(context.Background(), &severity)

		}
	}

	if err != nil {
		glog.Errorln(err)
		response.WriteHeaderAndEntity(http.StatusInternalServerError, &pb.SeverityResponse{Error: &pb.Error{Text: err.Error()}})
		return
	} else {
		response.WriteHeaderAndEntity(http.StatusOK, rsp)
	}

}
