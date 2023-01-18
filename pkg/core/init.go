package core

import (
	pb "github.com/uni-school/user-microservice/proto"
)

type Server struct {
	pb.UnimplementedUserServer
}
