package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	Name        string    `json:"Name" binding:"required"`
	Description string    `json:"Description" binding:"required"`
	Location    string    `json:"Location" binding:"required"`
	EventDate   time.Time `json:"EventDate" binding:"required"`
	UserId      int       `json:"UserId" binding:"required"`
	User        User      `gorm:"foreignKey:UserId" json:"User"`
	gorm.Model
}
