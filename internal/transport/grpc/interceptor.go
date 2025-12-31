package grpc

import (
	"context"
	logger2 "order-service/pkg/logger"

	"google.golang.org/grpc"
)

func InjectLoggerInterceptor(log *logger2.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = logger2.WithLogger(ctx, log)
		return handler(ctx, req)
	}
}
