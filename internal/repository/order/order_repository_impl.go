package order

import (
	"context"
	"database/sql"
	"evermosTest/internal/entity"
)

type orderRepositoryImpl struct {
	db *sql.DB
}

func (o *orderRepositoryImpl) Save(ctx context.Context, newOrder *entity.Order) (*entity.Order, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db *sql.DB) Repository {
	return &orderRepositoryImpl{db: db}
}
