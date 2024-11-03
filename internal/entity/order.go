package entity

import (
	"evermosTest/internal/handler/response"
	"time"
)

type Order struct {
	Id           int       `gorm:"primary_key;autoIncrement"`
	ProductId    int       `gorm:"column:product_id"`
	Quantity     int       `gorm:"column:quantity"`
	TotalPrice   int       `gorm:"column:total_price"`
	UpdateTime   time.Time `gorm:"column:update_time;autoCreateTime;autoUpdateTime;type:timestamp;default:NOW()"`
	CreationTime time.Time `gorm:"column:creation_time;autoCreateTime;<-:create;type:timestamp;default:NOW()"`
}

func (o *Order) ToResponse() *response.Order {
	return &response.Order{
		OrderId:    o.Id,
		ProductId:  o.ProductId,
		Quantity:   o.Quantity,
		TotalPrice: o.TotalPrice,
	}
}
