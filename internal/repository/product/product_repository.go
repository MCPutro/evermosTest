package product

import (
	"context"
	"database/sql"
	"evermosTest/internal/entity"
)

type Repository interface {
	Save(ctx context.Context, newProduct *entity.Product) (*entity.Product, error)
	FindById(ctx context.Context, ID int) (*entity.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]*entity.Product, error)
	Update(ctx context.Context, newProduct *entity.Product) (*entity.Product, error)
	Remove(ctx context.Context, ID int) error
}
