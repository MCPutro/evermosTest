package repository

import "evermosTest/internal/repository/product"

type Repository interface {
	ProductRepository() product.Repository
}
