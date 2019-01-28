package client

import (
	"fmt"
	"github.com/carmanzhang/ks-alert-client/pkg/option"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"k8s.io/klog/glog"
	"time"
)

var clientLBConn *grpc.ClientConn

func GetDispatcherGrpcLoadBalanceClient() (*grpc.ClientConn, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	target := fmt.Sprintf("%s:%s", *option.DispatcherServiceHost, *option.DispatcherServicePort)
	clientLBConn, err = grpc.DialContext(ctx, target, grpc.WithInsecure())

	if err != nil {
		glog.Errorln(err.Error())
	}

	return clientLBConn, err
}
