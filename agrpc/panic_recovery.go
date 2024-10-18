package agrpc

import (
	"context"
	"log"
	"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func PanicRecoveryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			// 打印 panic 信息和堆栈
			log.Printf("Panic recovered: %v\n%s", r, debug.Stack())
			// 返回 gRPC 内部错误
			err = status.Errorf(codes.Internal, "Internal server error")
		}
	}()

	// 调用实际的 gRPC handler
	return handler(ctx, req)
}
