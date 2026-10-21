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
	UserId      int       `json:"user_id"`
}

var events []Event = []Event{}

// AddEvent adds a new event to the events slice
func (e *Event) AddEvent() {
	events = append(events, *e)
}

// ShowEvents returns the list of all events
func ShowEvents() []Event {
	return events
}
