// models/book.go

package models

type Product struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	ProductName string `json:"productname"`
	Price       int    `json:"price"`
	//Deadline   time.Time `json:"deadline"`
	//CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	//UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
