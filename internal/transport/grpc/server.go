package grpc

import (
	"net"
	"order-service/pkg/api/test"

	"google.golang.org/grpc"
)

type Server struct {
	srv *grpc.Server
	h   Handler
}

func NewServer(h Handler) *Server {
	return &Server{
		srv: grpc.NewServer(),
		h:   h,
	}
}
func (s *Server) StartServer() error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}
	test.RegisterOrderServiceServer(s.srv, &s.h)
	return s.srv.Serve(lis)
}
