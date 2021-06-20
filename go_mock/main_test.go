package main

import (
	"testing"

	mock_main "github.com/DaisukeMatsumoto0925/go_mock/mock"
	"github.com/golang/mock/gomock"
)

// type ApiClientMock struct{}

// func (a *ApiClientMock) Request(data string) (string, error) {
// 	return data, nil
// }

func TestSample(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// d := &DataRegister{}
	// d.client = &ApiClientMock{}
	mockApiClient := mock_main.NewMockApiClient(ctrl)
	mockApiClient.EXPECT().Request("bar").Return("bar", nil)

	d := &DataRegister{}
	d.client = mockApiClient
	expected := "bar"

	res, err := d.Register(expected)
	if err != nil {
		t.Fatal("Regester error!", err)
	}
	if res != expected {
		t.Fatal("Value does not match.")
	}
}
