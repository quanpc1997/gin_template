package model

import (
	"time"
)

type Product struct {
	ID          uint      `gorm:"type:varchar(36);primaryKey"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"type:decimal(10,2);not null"`
	CreatedAt   time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`
}
