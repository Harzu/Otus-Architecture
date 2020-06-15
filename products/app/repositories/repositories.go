package repositories

import (
	"otus-products/app/repositories/products"
	"otus-products/app/system/database/psql"
)

type Repositories struct {
	productsRepo products.Repository
}

func New(client psql.Repository) *Repositories {
	return &Repositories{
		productsRepo: products.NewRepository(client),
	}
}

func (r *Repositories) GetProductsRepo() products.Repository {
	return r.productsRepo
}
