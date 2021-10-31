package order

import (
	"reflect"
	"testing"
)

func TestNewOrderService(t *testing.T) {
	type args struct {
		cfgs []OrderConfiguration
	}
	tests := []struct {
		name    string
		args    args
		want    *OrderService
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOrderService(tt.args.cfgs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOrderService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderService() = %v, want %v", got, tt.want)
			}
		})
	}
}
