package client

import (
	"fmt"
	"github.com/carmanzhang/ks-alert-client/pkg/constant"
	"github.com/carmanzhang/ks-alert/pkg/registry"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

var clientLBConn *grpc.ClientConn

func GetDispatcherGrpcLoadBalanceClient(address ...string) (*grpc.ClientConn, error) {
	var svc string
	var etcd string
	if address != nil && len(address) > 0 {
		svc = address[0]
		if len(address) > 1 {
			etcd = address[1]
		}
	} else {
		svc = constant.DispatcherServerName
		etcd = fmt.Sprintf("http://%s:%d", constant.ETCDHost, constant.ETCDPort)
	}

	if clientLBConn != nil {
		return clientLBConn, nil
	}

	r := registry.NewResolver(svc)
	b := grpc.RoundRobin(r)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	clientLBConn, err = grpc.DialContext(ctx, etcd, grpc.WithInsecure(), grpc.WithBalancer(b))
	return clientLBConn, err
}
