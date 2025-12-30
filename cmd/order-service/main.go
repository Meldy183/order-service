package main

import (
	"order-service/internal/repository"
	"order-service/internal/service"
	"order-service/internal/transport/grpc"
)

func main() {
	orderRepo := repository.NewOrderRepository()
	orderService := service.NewOrderService(&orderRepo)
	handler := grpc.NewHandler(orderService)
	server := grpc.NewServer(handler)
	err := server.StartServer()
	if err != nil {
		panic(err)
	}
}
