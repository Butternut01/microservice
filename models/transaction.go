package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model // Добавляет ID, CreatedAt, UpdatedAt, DeletedAt
	CartID     string  `gorm:"not null"`
	CustomerID string  `gorm:"not null"`
	Email      string  `gorm:"not null"`
	Status     string  `gorm:"not null;default:'pending'"` // "pending", "paid", "declined"
	TotalPrice float64
}
