package order

import (
	"context"
	"evermosTest/internal/entity"
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

func NewRepository(db *gorm.DB) Repository {
	return &orderRepositoryImpl{db: db}
}
