package product

import (
	"reflect"
	"testing"
)

func TestNewProduct(t *testing.T) {
	type args struct {
		name        string
		description string
		price       float64
	}
	tests := []struct {
		name    string
		args    args
		want    Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewProduct(tt.args.name, tt.args.description, tt.args.price)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
