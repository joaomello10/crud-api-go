package models

import (
	"time"
)

type Note struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Title     string
	Body      string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
