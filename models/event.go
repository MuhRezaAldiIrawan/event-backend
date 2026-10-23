package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	EventDate   time.Time `json:"event_date" binding:"required"`
	UserID      uint      `json:"user_id" binding:"required"`
	User        User      `gorm:"foreignKey:UserID;references:ID" json:"-" binding:"-"`
}
