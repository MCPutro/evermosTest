package product

import (
	"context"
	"evermosTest/internal/handler/request"
	"evermosTest/internal/handler/response"
)

type Service interface {
	Create(ctx context.Context, product *request.Product) (*response.Product, error)
	PriceAndStockAdjustment(ctx context.Context, productId int, isDeduct bool, n int, newPrice int) (*response.Product, error)
	Delete(ctx context.Context, productId int) error
	Checkout(ctx context.Context, product *request.Checkout) (*response.Order, error)
}
