package service

import (
	"context"
	"order-service/internal/models"
)

type OrderService struct {
	r OrderRepository
}

type OrderRepository interface {
	Select(ctx context.Context, id models.ID) (models.Order, error)
	Insert(ctx context.Context, order models.Order) (models.Order, error)
	Update(ctx context.Context, order models.Order) (models.Order, error)
	Delete(ctx context.Context, id models.ID) (bool, error)
	SelectAll(ctx context.Context) []models.Order
}

func NewOrderService(repository OrderRepository) OrderService {
	return OrderService{r: repository}
}

func (s *OrderService) GetOrder(ctx context.Context, id models.ID) (models.Order, error) {
	return s.r.Select(ctx, id)
}
func (s *OrderService) GetAllOrders(ctx context.Context) []models.Order {
	return s.r.SelectAll(ctx)
}
func (s *OrderService) CreateOrder(ctx context.Context, order models.Order) (models.Order, error) {
	return s.r.Insert(ctx, order)
}
func (s *OrderService) DeleteOrder(ctx context.Context, id models.ID) (bool, error) {
	return s.r.Delete(ctx, id)
}
func (s *OrderService) UpdateOrder(ctx context.Context, order models.Order) (models.Order, error) {
	return s.r.Update(ctx, order)
}
