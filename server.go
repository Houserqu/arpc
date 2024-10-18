package arpc

import (
	"github.com/Houserqu/arpc/agrpc"
	"github.com/Houserqu/arpc/ahttp"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type Server struct {
	GrpcServer *agrpc.GRPCServer
	HTTPServer *ahttp.HTTPServer
}

type ServerConfig struct {
	GrpcInterceptors []grpc.UnaryServerInterceptor
}

func NewServer(args ...any) *Server {
	InitConfig()
	InitMysql()
	InitRedis()

	// 服务配置
	var serverConfig ServerConfig
	if len(args) > 0 {
		serverConfig = args[0].(ServerConfig)
	}

	server := Server{}

	server.GrpcServer = agrpc.NewGrpcServer(serverConfig.GrpcInterceptors)
	server.HTTPServer = ahttp.NewHttpServer()

	return &server
}

// Start 启动服务
func (s *Server) Start() {
	go s.StartGrpc()
	go s.StartHttp()

	select {}
}

func (server *Server) StartGrpc() {
	server.GrpcServer.StartGrpcServer(viper.GetString("grpc.addr"))
}

func (server *Server) StartHttp() {
	server.HTTPServer.StartHttpServer(viper.GetString("http.addr"))
}
