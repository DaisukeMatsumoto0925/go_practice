package memory

import (
	"reflect"
	"sync"
	"testing"

	"github.com/DaisukeMatsumoto0925/tavern/domain/customer"
	"github.com/google/uuid"
)

func TestMemoryRepository_Add(t *testing.T) {
	type fields struct {
		customers map[uuid.UUID]customer.Customer
		Mutex     sync.Mutex
	}
	type args struct {
		c customer.Customer
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
			mr := &MemoryRepository{
				customers: tt.fields.customers,
				Mutex:     tt.fields.Mutex,
			}
			if err := mr.Add(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("MemoryRepository.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemoryRepository_Get(t *testing.T) {
	type fields struct {
		customers map[uuid.UUID]customer.Customer
		Mutex     sync.Mutex
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    customer.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := &MemoryRepository{
				customers: tt.fields.customers,
				Mutex:     tt.fields.Mutex,
			}
			got, err := mr.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemoryRepository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemoryRepository.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemoryRepository_Update(t *testing.T) {
	type fields struct {
		customers map[uuid.UUID]customer.Customer
		Mutex     sync.Mutex
	}
	type args struct {
		c customer.Customer
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
			mr := &MemoryRepository{
				customers: tt.fields.customers,
				Mutex:     tt.fields.Mutex,
			}
			if err := mr.Update(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("MemoryRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
