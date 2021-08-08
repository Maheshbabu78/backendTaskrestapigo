// controllers/books.go

package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"main.go/productCrud/models"
)

//create a product
func CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Please check the fields")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data inseted sucessfully": input})
}

// Get all products
func FindProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	db.Find(&products)
	c.JSON(http.StatusOK, gin.H{"List of Products": products})
}

// get product by id
func Findproductbyid(c *gin.Context) {
	var product models.Product
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		log.Println("The product is not found with this Id:", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Product": product})
}

// Update a product by id
func Updateproduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var product models.Product
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		log.Println("The product is not available with this Id", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found!"})
		return
	}
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&product).Updates(input)
	c.JSON(http.StatusOK, gin.H{"Updated Product Is": product})
}

// Delete a product by id
func DeleteProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var product models.Product
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		log.Println("Please check Id:", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product is not found!"})
		return
	}
	db.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"deleted Product is": product})
}
