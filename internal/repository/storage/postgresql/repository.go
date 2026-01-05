package postgresql

import (
	"context"
	"fmt"
	"order-service/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	pool    *pgxpool.Pool
	builder sq.StatementBuilderType
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		pool:    pool,
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r *OrderRepository) Select(ctx context.Context, id models.ID) (models.Order, error) {
	query := r.builder.Select("*").From("order").Where(sq.Eq{"id": id})
	sql, args, err := query.ToSql()
	if err != nil {
		return models.Order{}, fmt.Errorf("select order: %w", err)
	}
	var order models.Order
	err = r.pool.QueryRow(ctx, sql, args...).Scan(&order)
	if err != nil {
		return models.Order{}, fmt.Errorf("select order: %w", err)
	}
	return order, nil
}
func (r *OrderRepository) Insert(ctx context.Context, order models.Order) (models.Order, error) {
	newID := uuid.NewString()
	order.ID = newID
	query := r.builder.Insert("order").
		Columns("id", "item", "quantity").
		Values(order.ID, order.Item, order.Quantity)
	sql, args, err := query.ToSql()
	if err != nil {
		return models.Order{}, fmt.Errorf("insert order: %w", err)
	}
	_, err = r.pool.Exec(ctx, sql, args...)
	if err != nil {
		return models.Order{}, fmt.Errorf("insert order: %w", err)
	}
	return order, nil
}
func (r *OrderRepository) Update(ctx context.Context, order models.Order) (models.Order, error) {
	query := r.builder.Update("order").
		Set("quantity", order.Quantity).
		Set("item", order.Item).
		Where(sq.Eq{"id": order.ID})
	sql, args, err := query.ToSql()
	if err != nil {
		return models.Order{}, fmt.Errorf("update order: %w", err)
	}
	_, err = r.pool.Exec(ctx, sql, args...)
	if err != nil {
		return models.Order{}, fmt.Errorf("update order: %w", err)
	}
	return order, nil
}
func (r *OrderRepository) Delete(ctx context.Context, id models.ID) (bool, error) {
	query := r.builder.Delete("order").Where(sq.Eq{"id": id})
	sql, args, err := query.ToSql()
	if err != nil {
		return false, fmt.Errorf("delete order: %w", err)
	}
	_, err = r.pool.Exec(ctx, sql, args...)
	if err != nil {
		return false, fmt.Errorf("delete order: %w", err)
	}
	return true, nil
}
func (r *OrderRepository) SelectAll(ctx context.Context) []models.Order {
	query := r.builder.Select("*").From("order")
	sql, args, err := query.ToSql()
	if err != nil {
		return []models.Order{}
	}
	var orders []models.Order
	err = r.pool.QueryRow(ctx, sql, args...).Scan(&orders)
	if err != nil {
		return []models.Order{}
	}
	return orders
}
