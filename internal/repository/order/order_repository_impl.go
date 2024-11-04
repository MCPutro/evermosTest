package order

import (
	"context"
	"evermosTest/internal/entity"
	"fmt"
	"gorm.io/gorm"
)

type orderRepositoryImpl struct {
	db *gorm.DB
}

func (o *orderRepositoryImpl) Save(ctx context.Context, newOrder *entity.Order) (*entity.Order, error) {
	result := o.db.WithContext(ctx).Create(&newOrder)

	if result.Error != nil {
		return nil, result.Error
	}

	return newOrder, nil
}

func (o *orderRepositoryImpl) GetAll(ctx context.Context) (entity.OrderList, error) {
	var list entity.OrderList

	result := o.db.WithContext(ctx).Preload("Product").Find(&list)

	if result.Error != nil {
		return nil, result.Error
	}

	for _, order := range list {
		fmt.Println(order)
	}
	return list, nil

}

func NewRepository(db *gorm.DB) Repository {
	return &orderRepositoryImpl{db: db}
}
