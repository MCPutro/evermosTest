package order

import (
	"context"
	"evermosTest/internal/handler/response"
)

type Service interface {
	GetAllOrder(ctx context.Context) ([]*response.Order, error)
}
