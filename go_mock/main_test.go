package main

import "testing"

type ApiClientMock struct{}

func (a *ApiClientMock) Request(data string) (string, error) {
	return data, nil
}

func TestSample(t *testing.T) {
	d := &DataRegister{}
	d.client = &ApiClientMock{}
	expected := "bar"
	res, err := d.Register(expected)
	if err != nil {
		t.Fatal("Regester error!", err)
	}
	if res != expected {
		t.Fatal("Value does not match.")
	}
}
