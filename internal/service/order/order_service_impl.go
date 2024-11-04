package order

import (
	"context"
	"evermosTest/internal/handler/response"
	"evermosTest/internal/repository/order"
)

type ServiceImpl struct {
	order order.Repository
}

func (o *ServiceImpl) GetAllOrder(ctx context.Context) ([]*response.Order, error) {
	orderList, err := o.order.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return orderList.ToResponse(), nil
}

func NewService(order order.Repository) Service {
	return &ServiceImpl{order: order}
}
