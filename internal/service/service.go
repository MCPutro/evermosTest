package service

import (
	"evermosTest/internal/repository"
	"evermosTest/internal/service/order"
	"evermosTest/internal/service/product"
)

type Service interface {
	OrderService() order.Service
	ProductService() product.Service
}

type serviceManager struct {
	repoManager repository.Repository
}

func (s *serviceManager) OrderService() order.Service {
	return order.NewService(s.repoManager.OrderRepository())
}

func (s *serviceManager) ProductService() product.Service {
	return product.NewService(s.repoManager.ProductRepository(), s.repoManager.OrderRepository())
}

func NewServiceManager(repoManager repository.Repository) Service {
	return &serviceManager{repoManager: repoManager}
}
