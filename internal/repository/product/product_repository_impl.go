package product

import (
	"context"
	"database/sql"
	"evermosTest/internal/entity"
)

type productRepositoryImpl struct {
	db *sql.DB
}

func (p *productRepositoryImpl) Save(ctx context.Context, newProduct *entity.Product) (*entity.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productRepositoryImpl) FindById(ctx context.Context, ID int) (*entity.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]*entity.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productRepositoryImpl) Update(ctx context.Context, newProduct *entity.Product) (*entity.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *productRepositoryImpl) Remove(ctx context.Context, ID int) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db *sql.DB) Repository {
	return &productRepositoryImpl{db: db}
}
