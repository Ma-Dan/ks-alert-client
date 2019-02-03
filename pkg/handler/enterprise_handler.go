package handler

import (
	"context"
	"errors"
	"kubesphere.io/ks-alert-client/pkg/client"
	"kubesphere.io/ks-alert/pkg/pb"
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"net/http"
)

func CreateEnterprise(request *restful.Request, response *restful.Response) {
	var enp = pb.Enterprise{}
	err := request.ReadEntity(&enp)
	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	conn, err := client.GetDispatcherGrpcLoadBalanceClient()

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	cli := pb.NewEnterpriseHandlerClient(conn)
	rsp, err := cli.CreateEnterprise(context.Background(), &enp)

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	} else {
		response.WriteHeaderAndEntity(http.StatusOK, rsp)
	}

}
func RetrieveEnterprise(request *restful.Request, response *restful.Response) {
	entID := request.QueryParameter("enterprise_id")
	entName := request.QueryParameter("enterprise_name")

	conn, err := client.GetDispatcherGrpcLoadBalanceClient()

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	cli := pb.NewEnterpriseHandlerClient(conn)

	rsp, err := cli.GetEnterprise(context.Background(), &pb.EnterpriseSpec{
		EnterpriseId:   entID,
		EnterpriseName: entName,
	})

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	} else {
		response.WriteHeaderAndEntity(http.StatusOK, rsp)
	}

}
func UpdateEnterprise(request *restful.Request, response *restful.Response) {
	var ent = pb.Enterprise{}
	err := request.ReadEntity(&ent)
	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	if ent.EnterpriseId == "" && ent.EnterpriseName == "" {
		response.WriteHeaderAndEntity(http.StatusOK, errors.New("enterprise_id or enterprise_name must be specified"))
		return
	}

	conn, err := client.GetDispatcherGrpcLoadBalanceClient()

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	cli := pb.NewEnterpriseHandlerClient(conn)

	rsp, err := cli.UpdateEnterprise(context.Background(), &ent)

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	} else {
		response.WriteHeaderAndEntity(http.StatusOK, rsp)
	}
}

func DeleteEnterprise(request *restful.Request, response *restful.Response) {
	entID := request.QueryParameter("enterprise_id")
	entName := request.QueryParameter("enterprise_name")
	conn, err := client.GetDispatcherGrpcLoadBalanceClient()

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	}

	cli := pb.NewEnterpriseHandlerClient(conn)

	rsp, err := cli.DeleteEnterprise(context.Background(), &pb.EnterpriseSpec{
		EnterpriseId:   entID,
		EnterpriseName: entName,
	})

	if err != nil {
		glog.Errorln(err)
		response.WriteError(http.StatusInternalServerError, err)
		return
	} else {
		response.WriteHeaderAndEntity(http.StatusOK, rsp)
	}
}

//
//
//func HandlerError(err error, response *restful.Response) {
//	if err != nil {
//		glog.Errorln(err)
//		response.WriteError(http.StatusInternalServerError,err)
//		return
//	}
//}
