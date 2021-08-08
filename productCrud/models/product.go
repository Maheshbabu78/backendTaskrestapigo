// models/book.go

package models

import "github.com/jinzhu/gorm"

//product struct
type Product struct {
	gorm.Model
	ProductName string  `json:"productname" binding:"required"`
	Price       int     `json:"price" binding:"required"`
	Address     Address `gorm:"embedded;embeddedPrefix:address"`
}

//embedded struct
type Address struct {
	Village string `json:"village" binding:"required"`
	State   string `json:"state" binding:"required"`
	Pincode int    `json:"pincode" binding:"required"`
}
