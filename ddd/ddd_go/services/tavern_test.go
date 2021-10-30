package services

import (
	"testing"

	"github.com/google/uuid"
)

func TestTavern_Order(t *testing.T) {
	type fields struct {
		OrderService   *OrderService
		BillingService interface{}
	}
	type args struct {
		customer uuid.UUID
		products []uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tavern{
				OrderService:   tt.fields.OrderService,
				BillingService: tt.fields.BillingService,
			}
			if err := tr.Order(tt.args.customer, tt.args.products); (err != nil) != tt.wantErr {
				t.Errorf("Tavern.Order() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
