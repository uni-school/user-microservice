package custom

import "google.golang.org/grpc/codes"

type CustomGRPCError struct {
	StatusCode codes.Code
	Message    string
}

func (e *CustomGRPCError) Error() string {
	return e.Message
}

func NewCustomGRPCError(statusCode codes.Code, message string) *CustomGRPCError {
	return &CustomGRPCError{statusCode, message}
}
