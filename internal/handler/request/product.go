package request

import "evermosTest/internal/entity"

type Product struct {
	Name  string `json:"product_name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

func (p *Product) ToEntity() *entity.Product {
	return &entity.Product{
		Name:  p.Name,
		Price: p.Price,
		Stock: p.Stock,
	}
}

type ProductStockIn struct {
	ProductId int `json:"product_id"`
	AddStock  int `json:"add_stock"`
}

type ProductPriceAdjust struct {
	ProductId int `json:"product_id"`
	NewPrice  int `json:"new_price"`
}
