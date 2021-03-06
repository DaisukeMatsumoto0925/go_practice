package main

import (
	"github.com/DaisukeMatsumoto0925/tavern/domain/product"
	"github.com/DaisukeMatsumoto0925/tavern/services/order"
	servicetavern "github.com/DaisukeMatsumoto0925/tavern/services/tavern"
	"github.com/google/uuid"
)

func main() {

	products := productInventory()
	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository(""),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}

	tavern, err := servicetavern.NewTavern(
		servicetavern.WithOrderService(os))
	if err != nil {
		panic(err)
	}

	uid, err := os.AddCustomer("Percy")
	if err != nil {
		panic(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)
	if err != nil {
		panic(err)
	}

}

func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		panic(err)
	}
	peenuts, err := product.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	products := []product.Product{
		beer, peenuts, wine,
	}
	return products
}
