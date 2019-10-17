package models

import (
	"sync"

	"github.com/jinzhu/gorm"
)

type DatabaseInstance struct {
}

var instance *DatabaseInstance
var once sync.Once

var db *gorm.DB
var err error

func GetInstance() *DatabaseInstance {
	once.Do(func() {
		instance = &DatabaseInstance{}
	})
	return instance
}

func (d *DatabaseInstance) InitDatabase() (result bool, err error) {
	db, err = gorm.Open("mysql", "test:Test!pw1234@(localhost)/sudoku?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return false, err
	}
	// defer db.Close()
	db.SingularTable(true)
	return true, nil
}

func (d *DatabaseInstance) GetDatabase() *gorm.DB {
	return db
}

func init() {
	issucc, err := GetInstance().InitDatabase()
	if !issucc {
		panic(err)
	}
}
