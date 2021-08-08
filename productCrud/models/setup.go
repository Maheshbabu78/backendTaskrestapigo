package models

import (
	"fmt"
	"log"

	// go-sql-driver/mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SetupDB : initializing mysql database
func ConnecttoDb() *gorm.DB {
	USER := "root"
	PASS := "Mahesh@478"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "practice"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		log.Fatal(err.Error())

	}
	return db
}
