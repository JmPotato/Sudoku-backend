package models

import (
	"time"

	// Import go-sql-driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User model
type User struct {
	UID            uint32 `gorm:"primary_key"`
	Type           uint8
	Username       string
	Authentication string
	Score          uint32
	Submited       uint32
	Passed         uint32
	CreatedTime    time.Time
}

// CheckUserName checks whether the username exists.
func (user *User) CheckUserName(username string) (result bool) {
	db.Where("username = ?", username).First(user)
	if user.Username != username {
		return true
	}
	return false
}
