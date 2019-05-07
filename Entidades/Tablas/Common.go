package Tablas

import "time"

type Model struct {
	ID        string `gorm:"primary_key;uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Status bool
	DeletedAt *time.Time
}