package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/gorilla/mux"
	"github.com/xfpng345/go_micro_ser/products"
)

func main() {

	r := mux.NewRouter()
	r.Handle("/{name}", httptransport.NewServer(
		products.makeProductByNameEndpoint(products, products.NewService()),
		decodeProductByNameRequest,
		encodeResponse))
}

func decodeProductByNameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if name, ok := mux.Vars(r)["name"]; ok {
		return name, nil
	}
	return nil, errors.New("name is required")
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
