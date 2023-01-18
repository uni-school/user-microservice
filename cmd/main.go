package main

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/uni-school/user-microservice/libs/config"
	"github.com/uni-school/user-microservice/libs/constant"
	"github.com/uni-school/user-microservice/libs/util"
	"github.com/uni-school/user-microservice/pkg/core"
	pb "github.com/uni-school/user-microservice/proto"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func init() {
	config.ConfigApps()
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", util.GetPortApp()))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterUserServer(grpcServer, &core.Server{})
	if env := util.GetEnvironmentType(); env == constant.TEST {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}
