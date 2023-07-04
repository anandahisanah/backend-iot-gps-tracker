package models

import "time"

type Gps struct {
	Id       uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Datetime time.Time `gorm:"not null" json:"datetime"`
	Link     string     `gorm:"not null" json:"link"`
}
