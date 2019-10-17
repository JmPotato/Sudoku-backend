package models

import (
	"time"

	// Import go-sql-driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Puzzle model
type Puzzle struct {
	PID         uint32 `gorm:"primary_key"`
	Type        uint8
	Content     string
	Descriptor  string
	level       uint8
	Submited    uint32
	Passed      uint32
	CreatedTime time.Time
}
