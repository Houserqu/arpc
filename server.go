package arpc

import (
	"github.com/Houserqu/arpc/agrpc"
	"github.com/Houserqu/arpc/ahttp"
	"github.com/spf13/viper"
)

type Server struct {
	GrpcServer *agrpc.GRPCServer
	HTTPServer *ahttp.HTTPServer
}

func NewServer() *Server {
	// 初始化配置
	InitConfig()
	InitMysql()
	InitRedis()

	server := Server{}

	server.GrpcServer = agrpc.NewGrpcServer()
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
