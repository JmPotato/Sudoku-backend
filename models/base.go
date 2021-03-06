package models

import (
	"log"
	"sync"

	// Import go-sql-driver
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DatabaseInstance is a singleton pattern database instance.
type DatabaseInstance struct {
}

var instance *DatabaseInstance
var once sync.Once

var db *gorm.DB
var dbErr error

func getInstance() *DatabaseInstance {
	// Make sure a singleton pattern
	once.Do(func() {
		instance = &DatabaseInstance{}
	})
	return instance
}

func (d *DatabaseInstance) initDatabase() (result bool, err error) {
	db, dbErr = gorm.Open("mysql", "test:Test!pw1234@(localhost)/sudoku?charset=utf8&parseTime=True&loc=Local")
	if dbErr != nil {
		return false, dbErr
	}
	// defer db.Close()
	db.SingularTable(true)
	return true, nil
}

func init() {
	succ, err := getInstance().initDatabase()
	if !succ {
		log.Fatal(err)
	}
}
