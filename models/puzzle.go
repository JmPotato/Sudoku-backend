package models

import (
	"errors"
	"strings"
	"time"

	"github.com/JmPotato/Sudoku-backend/generator"
	"github.com/JmPotato/Sudoku-backend/utils"
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

//GetPuzzleByPID gets a puzzle by its PID
func (p *Puzzle) GetPuzzleByPID(pid uint32) error {
	return db.Where("pid = ?", pid).First(p).Error
}

//AddPuzzleByPID gets a puzzle by its PID
func (p *Puzzle) AddPuzzleByPID(pid uint32) error {
	db.Where("pid = ?", pid).First(p)
	if p.PID != 0 {
		return errors.New("puzzle already exists")
	}

	p.PID = pid
	p.Type = 1
	if p.Level == 0 {
		p.Level = 25
		p.Content = strings.TrimSpace(generator.Generate(25))
	} else {
		p.Content = strings.TrimSpace(generator.Generate(int(p.Level)))
	}
	p.Descriptor = utils.HashToMD5(strings.TrimSpace(p.Content))
	p.CreatedTime = time.Now()
	return db.Create(p).Error
}

// SavePuzzleByPID saves an user's info if it already exists.
func (p *Puzzle) SavePuzzleByPID(pid uint32) error {
	db.Where("pid = ?", pid).First(p)
	if p.PID != 0 {
		return errors.New("puzzle already exists")
	}
	return db.Save(p).Error
}
