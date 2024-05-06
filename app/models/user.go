package model

import "time"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Name     string
	Password string
	Birthday time.Time
}
