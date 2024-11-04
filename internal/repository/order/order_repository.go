package order

import (
	"context"
	"evermosTest/internal/entity"
)

type Repository interface {
	Save(ctx context.Context, newOrder *entity.Order) (*entity.Order, error)
	GetAll(ctx context.Context) (entity.OrderList, error)
}
