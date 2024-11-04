package response

type Order struct {
	OrderId     int    `json:"order_id"`
	ProductId   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	TotalPrice  int    `json:"total_price"`
	OrderDate   string `json:"order_date,omitempty"`
}
