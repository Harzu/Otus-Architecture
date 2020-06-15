package products

import (
	"fmt"
	"otus-products/app/entities"
	"otus-products/app/system/database/psql"

	"github.com/jmoiron/sqlx"
)

// Repository - interface for products repository
type Repository interface {
	GetProducts(productType string, limit uint64) (entities.Products, error)
}

type repositoryDB struct {
	dbClient *sqlx.DB
}

// NewRepository - create new products repository
func NewRepository(db psql.Repository) *repositoryDB {
	return &repositoryDB{dbClient: db.GetConnection()}
}

func (r *repositoryDB) GetProducts(productType string, limit uint64) (entities.Products, error) {
	query, args, err := prepareGetProductsQuery(productType, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare get products query: %w", err)
	}

	rows, err := r.dbClient.Queryx(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}

	var result entities.Products
	for rows.Next() {
		var product entities.Product
		if err = rows.StructScan(&product); err != nil {
			continue
		}

		result = append(result, &product)
	}

	return result, nil
}
