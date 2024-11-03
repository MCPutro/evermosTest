package entity

import (
	"evermosTest/internal/handler/response"
	"time"
)

type Product struct {
	Id           int       `gorm:"primary_key;autoIncrement;<-:create"`
	Name         string    `gorm:"column:name;size:100;NOT NULL"`
	Price        int       `gorm:"column:price;default:0"`
	Stock        int       `gorm:"column:stock;default:0"`
	UpdateTime   time.Time `gorm:"column:update_time;autoCreateTime;autoUpdateTime;type:timestamp;default:NOW()"`
	CreationTime time.Time `gorm:"column:creation_time;autoCreateTime;<-:create;type:timestamp;default:NOW()"`
	OrderId      Order     `gorm:"foreignKey:product_id;references:id"`
}

func (p *Product) ToResponse() *response.Product {
	return &response.Product{
		Id:    p.Id,
		Name:  p.Name,
		Price: p.Price,
		Stock: p.Stock,
	}
}
