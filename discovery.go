package arpc

import (
	"log"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		var t T
		return t
	}

	client := newServerClient(conn)
	clients[name] = client

	log.Println("Create client: ", name)

	return client
}
