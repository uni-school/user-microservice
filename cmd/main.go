package main

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/uni-school/user-microservice/shared/config"
	"github.com/uni-school/user-microservice/shared/constant"
	"github.com/uni-school/user-microservice/shared/runner"
	"github.com/uni-school/user-microservice/shared/util"
	"github.com/uni-school/user-microservice/pkg/core"
	pb "github.com/uni-school/user-microservice/proto"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func init() {
	config.ConfigApps()
	config.ConfigureDatabaseSQL(constant.POSTGRES)

	runner.RunSeeder()
}

func main() {
	var (
		opts []grpc.ServerOption
	)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", util.GetPortApp()))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(opts...)

	server := core.InitServer()

	pb.RegisterUserServiceServer(grpcServer, server)
	if env := util.GetEnvironmentType(); env == constant.TEST {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}
