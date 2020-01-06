package models

import (
	"errors"
	"time"
)

// User model
type User struct {
	UID            uint32    `gorm:"column:uid;primary_key" json:"uid"`
	Type           uint8     `gorm:"column:type" json:"type"`
	Username       string    `gorm:"column:username" json:"username"`
	Authentication string    `gorm:"column:authentication" json:"-"`
	Score          uint32    `gorm:"column:score" json:"score"`
	Submited       uint32    `gorm:"column:submited" json:"submited"`
	Passed         uint32    `gorm:"column:passed" json:"passed"`
	CreatedTime    time.Time `gorm:"column:created_time" json:"created_time"`
}

// GetUserByUsername gets a user's info by its username.
func (u *User) GetUserByUsername(username string) error {
	return db.Where("username = ?", username).First(u).Error
}

// GetUserByUID gets a user's info by its username.
func (u *User) GetUserByUID(uid uint32) error {
	return db.Where("uid = ?", uid).First(u).Error
}

// CreateUser creates a new user if it doesn't exist.
func (u *User) CreateUser() error {
	db.Where("username = ?", u.Username).First(u)
	if u.UID != 0 {
		return errors.New("user already exists")
	}
	u.CreatedTime = time.Now()
	return db.Create(u).Error
}

// DeleteUserByUsername deletes an user by its username if it already exists.
func (u *User) DeleteUserByUsername(username string) error {
	db.Where("username = ?", username).First(u)
	if u.UID == 0 {
		return errors.New("user doesn't exist")
	}
	return db.Delete(u).Error
}

// DeleteUserByUID deletes an user by its uid if it already exists.
func (u *User) DeleteUserByUID(uid uint32) error {
	db.Where("uid = ?", uid).First(u)
	if u.UID == 0 {
		return errors.New("user doesn't exist")
	}
	return db.Delete(u).Error
}

// SaveUserByUID saves an user's info if it already exists.
func (u *User) SaveUserByUID(uid uint32) error {
	if u.UID == 0 {
		return errors.New("user doesn't exist")
	}
	return db.Save(u).Error
}
