package agrpc

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/samber/lo"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// LoggingInterceptor 记录请求
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var printInfo string
	disableRpcs := viper.GetStringSlice("log.disable_req_rpc")
	if lo.Contains(disableRpcs, info.FullMethod) {
		printInfo = fmt.Sprintf("gRPC method: %s", info.FullMethod)
	} else {
		printInfo = fmt.Sprintf("gRPC method: %s, request: %v", info.FullMethod, req)
	}
	log.Println(strings.Join([]string{
		"level=INFO",
		"requestID=" + getRequestID(ctx),
		printInfo,
	}, " "))

	return handler(ctx, req)
}

// getRequestID 从上下文中获取 requestID
func getRequestID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if requestIDs, exists := md["request-id"]; exists {
			return requestIDs[0]
		}
	}
	// 生成request_id
	requestID := uuid.NewV4().String()
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("request-id", requestID))
	return requestID
}
