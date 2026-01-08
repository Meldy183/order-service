package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"order-service/internal/models"
	"time"

	"github.com/redis/go-redis/v9"
)

type OrderRepository interface {
	Select(ctx context.Context, id models.ID) (models.Order, error)
	Insert(ctx context.Context, order models.Order) (models.Order, error)
	Update(ctx context.Context, order models.Order) (models.Order, error)
	Delete(ctx context.Context, id models.ID) (bool, error)
	SelectAll(ctx context.Context) []models.Order
}

type CachedOrderRepository struct {
	repo OrderRepository
	rdb  *redis.Client
	ttl  time.Duration
}

func NewCachedOrderRepository(repo OrderRepository, rdb *redis.Client, ttl time.Duration) *CachedOrderRepository {
	return &CachedOrderRepository{
		repo: repo,
		rdb:  rdb,
		ttl:  ttl,
	}
}

func (c *CachedOrderRepository) orderKey(id models.ID) string {
	return fmt.Sprintf("order:%s", id)
}

func (c *CachedOrderRepository) Select(ctx context.Context, id models.ID) (models.Order, error) {
	key := c.orderKey(id)

	// Try cache first
	data, err := c.rdb.Get(ctx, key).Bytes()
	if err == nil {
		var order models.Order
		if err := json.Unmarshal(data, &order); err == nil {
			return order, nil
		}
	}

	// Cache miss - get from DB
	order, err := c.repo.Select(ctx, id)
	if err != nil {
		return models.Order{}, err
	}

	// Store in cache
	if data, err := json.Marshal(order); err == nil {
		_ = c.rdb.Set(ctx, key, data, c.ttl).Err()
	}

	return order, nil
}

func (c *CachedOrderRepository) Insert(ctx context.Context, order models.Order) (models.Order, error) {
	created, err := c.repo.Insert(ctx, order)
	if err != nil {
		return models.Order{}, err
	}

	// Cache new order
	key := c.orderKey(created.ID)
	if data, err := json.Marshal(created); err == nil {
		_ = c.rdb.Set(ctx, key, data, c.ttl).Err()
	}

	// Invalidate list cache
	_ = c.rdb.Del(ctx, "orders:all").Err()

	return created, nil
}

func (c *CachedOrderRepository) Update(ctx context.Context, order models.Order) (models.Order, error) {
	updated, err := c.repo.Update(ctx, order)
	if err != nil {
		return models.Order{}, err
	}

	// Update cache
	key := c.orderKey(updated.ID)
	if data, err := json.Marshal(updated); err == nil {
		_ = c.rdb.Set(ctx, key, data, c.ttl).Err()
	}

	// Invalidate list cache
	_ = c.rdb.Del(ctx, "orders:all").Err()

	return updated, nil
}

func (c *CachedOrderRepository) Delete(ctx context.Context, id models.ID) (bool, error) {
	ok, err := c.repo.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	// Remove from cache
	_ = c.rdb.Del(ctx, c.orderKey(id), "orders:all").Err()

	return ok, nil
}

func (c *CachedOrderRepository) SelectAll(ctx context.Context) []models.Order {
	key := "orders:all"

	// Try cache first
	data, err := c.rdb.Get(ctx, key).Bytes()
	if err == nil {
		var orders []models.Order
		if err := json.Unmarshal(data, &orders); err == nil {
			return orders
		}
	}

	// Cache miss - get from DB
	orders := c.repo.SelectAll(ctx)

	// Store in cache
	if data, err := json.Marshal(orders); err == nil {
		_ = c.rdb.Set(ctx, key, data, c.ttl).Err()
	}

	return orders
}
