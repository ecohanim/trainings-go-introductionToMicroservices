package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "tea",
		Price: 5.00,
		SKU: "dfg-abd-ghj",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
