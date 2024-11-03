package repository

import (
	"evermosTest/internal/repository/order"
	"evermosTest/internal/repository/product"
	"gorm.io/gorm"
	"sync"
)

var (
	productRepository     product.Repository
	productRepositoryOnce sync.Once

	orderRepository     order.Repository
	orderRepositoryOnce sync.Once
)

type Repository interface {
	ProductRepository() product.Repository
	OrderRepository() order.Repository
}

type repositoryManager struct {
	db *gorm.DB
}

func (r *repositoryManager) OrderRepository() order.Repository {
	orderRepositoryOnce.Do(func() {
		orderRepository = order.NewRepository(r.db)
	})
	return orderRepository
}

func (r *repositoryManager) ProductRepository() product.Repository {
	productRepositoryOnce.Do(func() {
		productRepository = product.NewRepository(r.db)
	})

	return productRepository
}

func NewRepositoryManager(db *gorm.DB) Repository {
	return &repositoryManager{db: db}
}
