package gateway

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"order-service/pkg/api/test"
)

type Gateway struct {
	httpServer *http.Server
	grpcConn   *grpc.ClientConn
}

func NewGateway() *Gateway {
	return &Gateway{}
}

func (g *Gateway) Start(ctx context.Context, httpAddr, grpcAddr string) error {
	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	g.grpcConn = conn

	mux := runtime.NewServeMux()
	err = test.RegisterOrderServiceHandler(ctx, mux, conn)
	if err != nil {
		return err
	}

	g.httpServer = &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	return g.httpServer.ListenAndServe()
}

func (g *Gateway) Stop(ctx context.Context) error {
	if g.httpServer != nil {
		if err := g.httpServer.Shutdown(ctx); err != nil {
			return err
		}
	}
	if g.grpcConn != nil {
		return g.grpcConn.Close()
	}
	return nil
}
