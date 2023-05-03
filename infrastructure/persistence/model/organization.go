package model

import "time"

type Organization struct {
	ID        string    `gorm:"primaryKey"`
	Name      string    `gorm:"unique"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
