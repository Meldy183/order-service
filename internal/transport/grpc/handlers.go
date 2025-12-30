package grpc

import (
	"context"
	"order-service/internal/models"
	"order-service/internal/service"
	"order-service/pkg/api/test"
)

type Handler struct {
	s service.OrderService
	test.UnimplementedOrderServiceServer
}

func NewHandler(s service.OrderService) Handler {
	return Handler{
		s: s,
	}
}
func (h *Handler) CreateOrder(ctx context.Context, request *test.CreateOrderRequest) (*test.CreateOrderResponse, error) {
	var order models.Order
	order.Quantity = int(request.Quantity)
	order.Item = request.Item
	respOrder, err := h.s.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	var response test.CreateOrderResponse
	response.Id = respOrder.ID
	return &response, nil
}
func (h *Handler) DeleteOrder(ctx context.Context, request *test.DeleteOrderRequest) (*test.DeleteOrderResponse, error) {
	var order models.Order
	order.ID = request.Id
	success, err := h.s.DeleteOrder(ctx, order.ID)
	if err != nil {
		return nil, err
	}
	var response test.DeleteOrderResponse
	response.Success = success
	return &response, nil
}
func (h *Handler) UpdateOrder(ctx context.Context, request *test.UpdateOrderRequest) (*test.UpdateOrderResponse, error) {
	var order models.Order
	order.ID = request.Id
	order.Quantity = int(request.Quantity)
	order.Item = request.Item
	_, err := h.s.UpdateOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	var response test.UpdateOrderResponse
	var orderDTO *test.Order
	orderDTO = &test.Order{
		Id:       request.Id,
		Item:     request.Item,
		Quantity: request.Quantity,
	}
	response.Order = orderDTO
	return &response, nil
}
func (h *Handler) GetOrder(ctx context.Context, request *test.GetOrderRequest) (*test.GetOrderResponse, error) {
	order, err := h.s.GetOrder(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	var response test.GetOrderResponse
	var orderDTO *test.Order
	orderDTO = &test.Order{
		Id:       order.ID,
		Item:     order.Item,
		Quantity: int32(order.Quantity),
	}
	response.Order = orderDTO
	return &response, nil
}
func (h *Handler) ListOrders(ctx context.Context, request *test.ListOrdersRequest) (*test.ListOrdersResponse, error) {
	var orderDTOs []*test.Order
	orders := h.s.GetAllOrders(ctx)
	for _, order := range orders {
		orderDTOs = append(orderDTOs, &test.Order{
			Id:       order.ID,
			Item:     order.Item,
			Quantity: int32(order.Quantity),
		})
	}
	response := &test.ListOrdersResponse{
		Orders: orderDTOs,
	}
	return response, nil
}
