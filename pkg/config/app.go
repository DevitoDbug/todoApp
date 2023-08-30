package config

import (
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func ConnectToDB() {
	DBConnectionString := "root:davi@tcp(localhost:3306)/bookstore?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", DBConnectionString)
	if err != nil {
		log.Printf("Could not connet to the DB:\n%v\n", err)
		return
	}
	db = d
}

func GetDBConnection() *gorm.DB {
	return db
}
