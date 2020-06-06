package repositories

import (
	"otus-auth/app/repositories/products"
	"otus-auth/app/system/database/psql"
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
