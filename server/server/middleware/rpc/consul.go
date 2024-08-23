package rpc

import (
	"api2db-server/config"
	"api2db-server/log"
	"api2db-server/pkg/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

var (
	SvrClient pb.ClientServiceClient
)

func NewSvrConn(svrName string) (*grpc.ClientConn, error) {
	consulInfo := config.GetGlobalConfig().ConsulConfig
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, svrName),
		// rpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		// rpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	if err != nil {
		log.Errorf("NewSvrConn with svrname %s err:%v", svrName, err)
		return nil, err
	}
	log.Info("NewSvrConn success")
	return conn, nil
}

func GetSvrClient() pb.ClientServiceClient {
	return SvrClient
}

func NewSvrClient(svrName string) pb.ClientServiceClient {
	conn, err := NewSvrConn(svrName)
	if err != nil {
		return nil
	}
	return pb.NewClientServiceClient(conn)
}

func InitSvrConn() {
	SvrClient = NewSvrClient(config.GetGlobalConfig().SvrConfig.SvrName)
}
