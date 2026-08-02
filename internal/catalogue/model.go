package catalogue

import (
	"time"
)

type Catalogue struct {
	ID                  uint      `json:"id" gorm:"primaryKey"`
	MaterialType        string    `json:"material_type" binding:"required"`
	PricePerKg          float64   `json:"price_per_kg" binding:"required"`
	LastUpdateDate      time.Time `json:"last_update_date"`
	MaterialDescription string    `json:"material_description" binding:"required"`
	CreatedAt           time.Time `json:"created_at"`
}
