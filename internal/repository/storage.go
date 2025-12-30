package repository

import (
	"context"
	"errors"
	"order-service/internal/models"
	"sync"

	"github.com/google/uuid"
)

type OrderRepository struct {
	db map[models.ID]models.Order
	mu sync.RWMutex
}

func NewOrderRepository() OrderRepository {
	return OrderRepository{
		make(map[models.ID]models.Order),
		sync.RWMutex{},
	}
}
func (r *OrderRepository) Select(ctx context.Context, id models.ID) (models.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if _, ok := r.db[id]; !ok {
		return models.Order{}, errors.New("order not found")
	}
	order := r.db[id]
	return order, nil
}
func (r *OrderRepository) Insert(ctx context.Context, order models.Order) (models.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if !ValidateOrder(order) {
		return models.Order{}, errors.New("invalid order")
	}
	id := uuid.NewString()
	r.db[id] = order
	order.ID = id
	return order, nil
}
func (r *OrderRepository) Update(ctx context.Context, order models.Order) (models.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.db[order.ID]; !ok {
		return models.Order{}, errors.New("order not found")
	}
	if !ValidateOrder(order) {
		return models.Order{}, errors.New("invalid order")
	}
	r.db[order.ID] = order
	return order, nil
}
func (r *OrderRepository) Delete(ctx context.Context, id models.ID) (bool, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.db[id]; !ok {
		return false, errors.New("order not found")
	}
	delete(r.db, id)
	return true, nil
}
func (r *OrderRepository) SelectAll(ctx context.Context) []models.Order {
	r.mu.RLock()
	defer r.mu.RUnlock()
	orders := make([]models.Order, 0, len(r.db))
	for _, order := range r.db {
		orders = append(orders, order)
	}
	return orders
}

func ValidateOrder(order models.Order) bool {
	if order.Item == "" {
		return false
	}
	if order.Quantity <= 0 {
		return false
	}
	return true
}
