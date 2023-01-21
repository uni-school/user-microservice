package custom

import (
	"encoding/json"
	"runtime"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CustomGRPCErrorHandler(err error) error {
	logrus.Infof("error: %#v\n", err)
	ge, ok := err.(*CustomGRPCError)
	if ok {
		return status.Error(ge.StatusCode, ge.Error())
	}
	me, ok := err.(*json.MarshalerError)
	if ok {
		return status.Error(codes.InvalidArgument, me.Error())
	}
	re, ok := err.(runtime.Error)
	if ok {
		return status.Error(codes.InvalidArgument, re.Error())
	}
	return status.Error(codes.Internal, err.Error())
}
