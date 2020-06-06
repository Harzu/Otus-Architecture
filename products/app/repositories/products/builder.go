package products

import (
	"github.com/Masterminds/squirrel"
)

const tableProducts = "products"

func prepareGetProductsQuery(productType string, limit uint64) (string, []interface{}, error) {
	psqlSq := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	rawRequest := psqlSq.Select("type", "name", "price").
		From(tableProducts).
		Where(squirrel.Eq{"type": productType}).
		Limit(limit)

	return rawRequest.ToSql()
}
