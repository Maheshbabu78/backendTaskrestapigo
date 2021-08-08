package main

import (
	"fmt"
	"testing"

	"main.go/productCrud/models"
)

func TestProductStruct(t *testing.T) {
	type test struct {
		input models.Product
		//input2     models.Address
		expected models.Product
	}
	product1 := models.Product{
		ProductName: "Samsung",
		Price:       20000,
	}
	product2 := models.Product{
		ProductName: "Apple",
		Price:       20000,
	}
	address1 := models.Address{
		Village: "bang",
		State:   "Karnataka",
		Pincode: 506344,
	}
	address2 := models.Address{
		Village: "Hyd",
		State:   "ts",
		Pincode: 506344,
	}
	areaTests := []struct {
		product models.Product
		address models.Address
	}{
		{product: product1, address: address1},
		{product: product2, address: address2},
	}
	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.product.ProductName, func(t *testing.T) {
			got := tt.product
			if got != tt.product {
				fmt.Println(tt.product)
			}
		})

	}
	for _, tt := range areaTests {

		t.Run(tt.address.State, func(t *testing.T) {
			got := tt.address
			if got != tt.address {
				fmt.Println(tt.address)
			}
		})

	}

}
