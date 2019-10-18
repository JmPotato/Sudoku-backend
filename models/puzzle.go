package models

import (
	"time"

	// Import go-sql-driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Puzzle model
type Puzzle struct {
	PID         uint32    `gorm:"primary_key" json:"pid"`
	Type        uint8     `json:"type"`
	Content     string    `json:"content"`
	Descriptor  string    `json:"descriptor"`
	level       uint8     `json:"level"`
	Submited    uint32    `json:"submited"`
	Passed      uint32    `json:"passed"`
	CreatedTime time.Time `json:"created_time"`
}
