package model

import "time"

type Robot struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Length    int       `json:"length"`
	Width     int       `json:"width"`
	Filename  string    `gorm:"not null" json:"filename"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
