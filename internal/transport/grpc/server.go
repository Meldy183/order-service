package grpc

import (
	"net"
	"order-service/pkg/api/test"
	"order-service/pkg/logger"

	"google.golang.org/grpc"
)

type Server struct {
	srv *grpc.Server
	h   Handler
}

func NewServer(h Handler, logger *logger.Logger) *Server {
	return &Server{
		srv: grpc.NewServer(
			grpc.UnaryInterceptor(InjectLoggerInterceptor(logger))),
		h: h,
	}
}
func (s *Server) StartServer(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	test.RegisterOrderServiceServer(s.srv, &s.h)
	return s.srv.Serve(lis)
}

func (s *Server) StopServer() error {
	s.srv.GracefulStop()
	return nil
}
