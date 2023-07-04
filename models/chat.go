package models

import "time"

type Chat struct {
	Id        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"not null" json:"username"`
	Message   string    `gorm:"not null" json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
