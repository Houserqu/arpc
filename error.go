package arpc

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BizError 业务错误
func BizError(message string, a ...any) error {
	return status.Errorf(codes.Internal, message, a...)
}
