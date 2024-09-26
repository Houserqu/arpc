package arpc

import (
	"github.com/Houserqu/arpc/agrpc"
	"github.com/Houserqu/arpc/ahttp"
)

type Server struct {
	GrpcServer *agrpc.GRPCServer
	HTTPServer *ahttp.HTTPServer
}

func NewServer() *Server {
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
	server.GrpcServer.StartGrpcServer(":8000")
}

func (server *Server) StartHttp() {
	server.HTTPServer.StartHttpServer(":8080")
}
