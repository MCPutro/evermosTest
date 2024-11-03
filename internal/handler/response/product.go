package response

type Product struct {
	Id    int    `json:"product_id"`
	Name  string `json:"product_name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}
