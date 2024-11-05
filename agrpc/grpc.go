package agrpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	Server *grpc.Server
}

// NewServer 创建新的 gRPC 服务并附加拦截器
func NewGrpcServer(userInterceptors []grpc.UnaryServerInterceptor) *GRPCServer {
	// 系统内置的拦截器
	interceptors := []grpc.UnaryServerInterceptor{
		validationInterceptor,    // 参数校验拦截器
		PanicRecoveryInterceptor, // panic 恢复拦截器
	}

	// 将用户自定义的拦截器和系统内置的拦截器合并
	interceptors = append(interceptors, userInterceptors...)

	chainUnaryInterceptor := grpc.ChainUnaryInterceptor(interceptors...)

	opts := []grpc.ServerOption{
		chainUnaryInterceptor,
		grpc.MaxRecvMsgSize(100 * 1024 * 1024),
	}

	s := grpc.NewServer(opts...)
	return &GRPCServer{Server: s}
}

// Start 启动 gRPC 服务
func (s *GRPCServer) StartGrpcServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Starting gRPC server on %s", addr)
	if err := s.Server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// Stop 关闭服务
func (s *GRPCServer) StopGrpcServer() {
	s.Server.GracefulStop()
}
