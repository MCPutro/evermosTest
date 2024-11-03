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
