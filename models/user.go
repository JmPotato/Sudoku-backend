package models

import (
	"time"

	// Import go-sql-driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User model
type User struct {
	UID            uint32    `gorm:"primary_key" json:"uid"`
	Type           uint8     `json:"type"`
	Username       string    `json:"username"`
	Authentication string    `json:"-"`
	Score          uint32    `json:"scroe"`
	Submited       uint32    `json:"submited"`
	Passed         uint32    `json:"passed"`
	CreatedTime    time.Time `json:"created_time"`
}

// CheckUserName checks whether the username exists.
func (user *User) CheckUserName() error {
	return db.Where("username = ?", user.Username).First(user).Error
}
