package service

import (
	"context"
	"evermosTest/internal/entity"
	"evermosTest/internal/handler/request"
	"evermosTest/internal/handler/response"
)

type ProductService interface {
	Create(ctx context.Context, product *entity.Product) (*response.Product, error)
	PriceAndStockAdjustment(ctx context.Context, productId int, isDeduct bool, n int, newPrice int) (*response.Product, error)
	Delete(ctx context.Context, productId int) error
	Checkout(ctx context.Context, product *request.Checkout) (*response.Order, error)
}
