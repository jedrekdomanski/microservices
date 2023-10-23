package data

import (
	"testing"
)

func TestProductValidation(t *testing.T) {
	p := &Product{
		Name: "Gopuccino",
		Price: 1.00,
		SKU: "abc-cde-efg",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
