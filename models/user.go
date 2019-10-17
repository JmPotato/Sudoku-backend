package models

import (
	"github.com/jinzhu/gorm"

	// Import go-sql-driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func connectDatabase() {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database.")
	}
	defer db.Close()
	// To-do: initilize the User table.
}
