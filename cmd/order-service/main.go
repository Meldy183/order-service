package main

import (
	"context"
	"fmt"
	"net/http"
	"order-service/internal/config"
	"order-service/internal/repository/storage/postgresql"
	"order-service/internal/service"
	"order-service/internal/transport/gateway"
	"order-service/internal/transport/grpc"
	"order-service/pkg/db"
	"order-service/pkg/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	cfg, err := config.MustParseConfig(configPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg, configPath)
	ctx := context.Background()
	log, err := logger.NewLogger(cfg.Server.Env)
	if err != nil {
		panic(err)
	}
	log.Info(ctx, "Logger initialized")
	pool, err := db.NewDataBase(cfg.DB)
	if err != nil {
		panic(err)
	}
	orderRepo := postgresql.NewOrderRepository(pool.Pool)
	orderService := service.NewOrderService(orderRepo)
	handler := grpc.NewHandler(orderService)
	server := grpc.NewServer(handler, log)

	wg := &sync.WaitGroup{}

	// Start gRPC server
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Info(ctx, "Starting gRPC server on port "+cfg.Server.Port)
		if err := server.StartServer(":" + cfg.Server.Port); err != nil {
			panic(err)
		}
	}()

	// Start HTTP gateway
	gw := gateway.NewGateway()
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Info(ctx, "Starting HTTP gateway on port "+cfg.Server.GatewayPort)
		if err := gw.Start(ctx, ":"+cfg.Server.GatewayPort, "localhost:"+cfg.Server.Port); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-gracefulStop

	log.Info(ctx, "Shutting down servers...")
	if err = gw.Stop(ctx); err != nil {
		log.Error(ctx, err.Error())
	}
	if err = server.StopServer(); err != nil {
		log.Error(ctx, err.Error())
	}
	log.Info(ctx, "Servers stopped")
}
