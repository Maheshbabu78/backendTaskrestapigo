// models/book.go

package models

type Product struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	ProductName string `json:"productname"`
	Price       int    `json:"price"`
}
