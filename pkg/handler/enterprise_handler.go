package handler

import (
	"github.com/emicklei/go-restful"
	"github.com/golang/glog"
	"kubesphere.io/ks-alert/pkg/dispatcher/pb"
)

func CreateEnterprise(request *restful.Request, response *restful.Response) {
	var enp = pb.Enterprise{}
	err := request.ReadEntity(&enp)
	if err != nil {
		glog.Errorln(err)
	}

	//enp.

}
func RetrieveEnterprise(request *restful.Request, response *restful.Response) {

}
func UpdateEnterprise(request *restful.Request, response *restful.Response) {

}
func DeleteEnterprise(request *restful.Request, response *restful.Response) {

}
