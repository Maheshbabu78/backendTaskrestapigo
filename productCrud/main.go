package main

import (
	"main.go/productCrud/models"
	"main.go/productCrud/routes"
)

func main() {

	db := models.ConnecttoDb()
	db.AutoMigrate(&models.Product{})
	r := routes.RoutersNavigation(db)
	r.Run()
}
