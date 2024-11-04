package entity

import (
	"evermosTest/internal/handler/response"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	Id           int            `gorm:"primary_key;autoIncrement;<-:create"`
	Name         string         `gorm:"column:name;size:100;NOT NULL"`
	Price        int            `gorm:"column:price;default:0"`
	Stock        int            `gorm:"column:stock;default:0"`
	UpdateTime   time.Time      `gorm:"column:update_time;autoCreateTime;autoUpdateTime;type:timestamp;default:NOW()"`
	CreationTime time.Time      `gorm:"column:creation_time;autoCreateTime;<-:create;type:timestamp;default:NOW()"`
	DeletedTime  gorm.DeletedAt `gorm:"column:delete_time"`
	OrderId      []Order        `gorm:"foreignKey:product_id;references:id"`
}

func (p *Product) ToResponse() *response.Product {
	return &response.Product{
		Id:    p.Id,
		Name:  p.Name,
		Price: p.Price,
		Stock: p.Stock,
	}
}

type ProductList []*Product

func (pl *ProductList) ToResponseProductList() []*response.Product {
	var tmp []*response.Product
	for _, product := range *pl {
		tmp = append(tmp, product.ToResponse())
	}

	return tmp
}
