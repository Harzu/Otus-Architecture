package entities

// Products - entity for list product
type Products []*Product

// Product - entity for product
type Product struct {
	Type  string `db:"type",json:"type"`
	Name  string `db:"name",json:"name"`
	Price int64  `db:"price",json:"price"`
}

// ToMap - convert struct to map
func (p Product) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"type":  p.Type,
		"name":  p.Name,
		"price": p.Price,
	}
}
