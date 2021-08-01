package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"main.go/productCrud/controllers"
)

func RoutersNavigation(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/products", controllers.FindProduct)
	r.POST("/createproduct", controllers.CreateProduct)
	r.GET("/product/:id", controllers.Findproductbyid)
	r.PATCH("/updateproduct/:id", controllers.Updateproduct)
	r.DELETE("deleteproduct/:id", controllers.DeleteProduct)
	return r
}
