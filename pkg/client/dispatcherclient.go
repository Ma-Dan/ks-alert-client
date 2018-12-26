package client

import (
	"github.com/carmanzhang/ks-alert/pkg/registry"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

var clientLBConn *grpc.ClientConn

func GetDispatcherGrpcLoadBalancerClient(svc string, etcdAddress string) (*grpc.ClientConn, error) {
	if clientLBConn != nil {
		return clientLBConn, nil
	}

	r := registry.NewResolver(svc)
	b := grpc.RoundRobin(r)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	clientLBConn, err = grpc.DialContext(ctx, etcdAddress, grpc.WithInsecure(), grpc.WithBalancer(b))
	return clientLBConn, err
}
