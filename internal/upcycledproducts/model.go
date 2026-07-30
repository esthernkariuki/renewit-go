package upcycledproducts

import (
	"renewit-go/internal/users"
	"time"
)

type Upcycledproduct struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	Image      string     `json:"image" binding:"required"`
	Quantity   int        `json:"quantity" binding:"required,min=1"`
	Type       string     `json:"type" binding:"required"`
	Price      float64    `json:"price" binding:"required"`
	UpdatedAt  time.Time  `json:"updated_at"`
	UpcyclerID uint       `json:"upcycler_id" binding:"required"`
	Upcycler   users.User `json:"upcycler" gorm:"foreignKey:UpcyclerID"`
}
