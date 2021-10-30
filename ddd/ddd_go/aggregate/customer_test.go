package aggregate

import (
	"testing"
)

func TestNewCustomer(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		// want    Customer
		wantErr bool
	}{
		{
			name: "Empty Name validation",
			args: args{
				name: "",
			},
			// want:    Customer{},
			wantErr: true,
		},
		{
			name: "Valid Name",
			args: args{
				name: "Daisuke Mt",
			},
			// want: Customer{
			// 	person:       &entity.Person{},
			// 	products:     []*entity.Item{},
			// 	transactions: []valueobject.Transaction{},
			// },
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCustomer(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("NewCustomer() = %v, want %v", got, tt.want)
			// }
		})
	}
}
