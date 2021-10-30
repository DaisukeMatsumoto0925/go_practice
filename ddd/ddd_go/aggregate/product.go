package aggregate

import (
	"errors"

	"github.com/DaisukeMatsumoto0925/ddd_go/entity"
	"github.com/google/uuid"
)

var (
	ErrMissingValues = errors.New("missing values")
)

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
}

func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() uuid.UUID {
	return p.item.ID
}

func (p Product) GetPrice() float64 {
	return p.price
}
