package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"main.go/productCrud/controllers"
	"main.go/productCrud/models"
	"main.go/productCrud/routes"
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
func TestGetProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := controllers.Findproductbyid
	router := gin.Default()
	router.GET("/product/:id", handler)

	req, err := http.NewRequest("GET", "/product/2", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	//assert.Equal(t, resp.Code, 200)
	assert.Equal(t, resp.Code, 200)
}
func TestGetProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := models.ConnecttoDb()
	db.AutoMigrate(&models.Product{})
	testRouter := routes.RoutersNavigation(db)
	req, err := http.NewRequest(http.MethodGet, "/product/8", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	w := httptest.NewRecorder()
	testRouter.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
