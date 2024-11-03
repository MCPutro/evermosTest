package request

type Checkout struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
