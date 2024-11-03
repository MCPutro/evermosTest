package entity

type Order struct {
	Id         int `json:"id"`
	ProductId  int `json:"product_id"`
	Quantity   int `json:"quantity"`
	TotalPrice int `json:"total_price"`
}
