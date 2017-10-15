package models

import (
	//"database/sql"
	"time"
)

// Event - MySQL table
type Event struct {
	ID                uint       `gorm:"primary_key" json:"event_id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
	CameraID          string     `gorm:"size:100;not null" json:"camera_id"`
	Prediction        string     `gorm:"size:100;not null" json:"prediction"`
	Thumbnail         string     `gorm:"size:512" json:"thumbnail"`
	StartingTimestamp uint       `gorm:"not null;default:0" json:"starting_timestamp"`
	IsViewed          bool       `gorm:"default:0" json:"is_viewed"`
}
