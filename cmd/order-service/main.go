package main

import (
	"context"
	"fmt"
	"order-service/internal/config"
	"order-service/internal/repository"
	"order-service/internal/service"
	"order-service/internal/transport/grpc"
	"order-service/pkg/logger"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./config/config.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	configPath := os.Getenv("CONFIG_PATH")
	cfg, err := config.MustParseConfig(configPath)
	fmt.Println(cfg)
	ctx := context.Background()
	logger, err := logger.NewLogger(cfg.Server.Env)
	if err != nil {
		panic(err)
	}
	logger.Info(ctx, "Logger initialized")
	orderRepo := repository.NewOrderRepository()
	orderService := service.NewOrderService(&orderRepo)
	handler := grpc.NewHandler(orderService)
	server := grpc.NewServer(handler, logger)
	err = server.StartServer(":" + cfg.Server.Port)
	if err != nil {
		panic(err)
	}
}
