package products

import (
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"otus-products/app/entities"
	"otus-products/app/generated/models"
	"otus-products/app/generated/restapi/operations"
	"otus-products/app/repositories/products"
	redisRepo "otus-products/app/system/database/redis"
	"time"

	"github.com/go-redis/redis"

	"github.com/go-openapi/runtime/middleware"
)

const cacheTTL = 5 * time.Second

// ProductsHandler - handler for search products
type ProductsHandler struct {
	Repository products.Repository
	RedisRepo  redisRepo.Repository
}

// NewHandler - create ProductsHandler
func NewHandler(repo products.Repository, redisRepository redisRepo.Repository) ProductsHandler {
	return ProductsHandler{Repository: repo, RedisRepo: redisRepository}
}

// GetProducts - search and return products by params
func (ph ProductsHandler) GetProducts(params operations.GetProductsParams) middleware.Responder {
	key := fmt.Sprintf("%s_%d", *params.Body.Type, *params.Body.Limit)
	md5Key := fmt.Sprintf("%x", md5.Sum([]byte(key)))

	if params.Etag != nil && *params.Etag == md5Key {
		res, err := ph.RedisRepo.GetConnection().Get(*params.Etag).Result()
		if err != nil && err != redis.Nil {
			errorMessage := err.Error()
			return operations.NewGetProductsInternalServerError().WithPayload(&models.FailResponse{
				Error: &errorMessage,
			})
		}

		if err == nil {
			var products entities.Products
			if err = json.Unmarshal([]byte(res), &products); err != nil {
				errorMessage := err.Error()
				return operations.NewGetProductsInternalServerError().WithPayload(&models.FailResponse{
					Error: &errorMessage,
				})
			}

			return operations.NewGetProductsOK().WithPayload(buildProductsModel(products))
		}
	}

	products, err := ph.Repository.GetProducts(*params.Body.Type, uint64(*params.Body.Limit))
	if err != nil {
		errorMessage := err.Error()
		if errors.Is(err, sql.ErrNoRows) {
			return operations.NewGetProductsBadRequest().WithPayload(&models.FailResponse{
				Error: &errorMessage,
			})
		}

		return operations.NewGetProductsInternalServerError().WithPayload(&models.FailResponse{
			Error: &errorMessage,
		})
	}

	time.Sleep(1 * time.Second)

	result := buildProductsModel(products)
	response := operations.NewGetProductsOK()

	if ph.RedisRepo.Enabled() {
		jsonRes, err := json.Marshal(result)
		if err != nil {
			errorMessage := "invalid json marshal"
			return operations.NewGetProductsInternalServerError().WithPayload(&models.FailResponse{
				Error: &errorMessage,
			})
		}

		if err := ph.RedisRepo.GetConnection().Set(md5Key, jsonRes, cacheTTL).Err(); err != nil {
			errorMessage := "invalid redis save"
			return operations.NewGetProductsInternalServerError().WithPayload(&models.FailResponse{
				Error: &errorMessage,
			})
		}

		response.SetETag(md5Key)
	}

	return response.WithPayload(result)
}

func buildProductsModel(products entities.Products) []*models.Products {
	result := make([]*models.Products, 0, len(products))
	for _, p := range products {
		item := models.Products{
			Name:  &p.Name,
			Type:  &p.Type,
			Price: &p.Price,
		}

		result = append(result, &item)
	}

	return result
}
