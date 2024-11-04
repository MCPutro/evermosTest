package entity

import (
	"evermosTest/internal/handler/response"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	Id           int            `gorm:"primary_key;autoIncrement"`
	ProductId    int            `gorm:"column:product_id"`
	Quantity     int            `gorm:"column:quantity"`
	TotalPrice   int            `gorm:"column:total_price"`
	UpdateTime   time.Time      `gorm:"column:update_time;autoCreateTime;autoUpdateTime;type:timestamp;default:NOW()"`
	CreationTime time.Time      `gorm:"column:creation_time;autoCreateTime;<-:create;type:timestamp;default:NOW()"`
	DeletedTime  gorm.DeletedAt `gorm:"column:delete_time"`
	Product      Product        `gorm:"foreignKey:product_id;references:id"` // belongs to
}

func (o *Order) ToResponse() *response.Order {
	return &response.Order{
		OrderId:     o.Id,
		ProductId:   o.ProductId,
		ProductName: o.Product.Name,
		Quantity:    o.Quantity,
		TotalPrice:  o.TotalPrice,
		OrderDate:   o.CreationTime.Format("2006-01-02 15:04:05"),
	}
}

type OrderList []*Order

func (ol *OrderList) ToResponse() []*response.Order {
	var temp []*response.Order
	for _, o := range *ol {
		temp = append(temp, o.ToResponse())
	}

	return temp
}
