package models

import (
	"time"
)

// Puzzle model
type Puzzle struct {
	PID         uint32    `gorm:"column:pid;primary_key" json:"pid"`
	Type        uint8     `gorm:"column:type" json:"type"`
	Content     string    `gorm:"column:content" json:"content"`
	Descriptor  string    `gorm:"column:descriptor" json:"descriptor"`
	Level       uint8     `gorm:"column:level" json:"level"`
	Submited    uint32    `gorm:"column:submited" json:"submited"`
	Passed      uint32    `gorm:"column:passed" json:"passed"`
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"`
}
