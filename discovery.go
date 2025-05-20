package arpc

import (
	"context"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// 存储服务的客户端
var clients = make(map[string]interface{})

// GetServerClient 获取服务客户端
func GetServerClient[T any](name string, newServerClient func(grpc.ClientConnInterface) T) T {
	// 如果客户端已经存在，则直接返回
	if client, ok := clients[name]; ok {
		return client.(T)
	}

	// 获取服务地址（优先取配置文件的）
	addr := viper.GetString("discovery." + name)
	if addr == "" {
		addr = name + ":8000"
	}

	// 创建客户端
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(viper.GetInt("grpc.max_msg_size")*1024*1024)),
		grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			requestID, ok := ctx.Value("request-id").(string)
			if !ok || requestID == "" {
				requestID = uuid.NewV4().String()
			}

			ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("request-id", requestID))

			return invoker(ctx, method, req, reply, cc)
		}),
	)
	if err != nil {
		var t T
		return t
	}

	client := newServerClient(conn)
	clients[name] = client

	log.Println("Create client: ", name)

	return client
}
