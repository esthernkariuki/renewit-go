package upcycledproducts

import (
	"renewit-go/internal/users"
	"time"
)

type Upcycledproduct struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	UpcycledClothes string     `json:"upcycled_clothes" form:"upcycled_clothes" binding:"required"`
	Description     string     `json:"description" form:"description"`
	Image           string     `json:"image" form:"image"`
	Quantity        int        `json:"quantity" form:"quantity" binding:"required,min=1"`
	Type            string     `json:"type" form:"type" binding:"required"`
	Material        string     `json:"material" form:"material"`
	Size            string     `json:"size" form:"size"`
	Color           string     `json:"color" form:"color"`
	Condition       string     `json:"condition" form:"condition"`
	Location        string     `json:"location" form:"location"`
	Price           float64    `json:"price" form:"price" binding:"required"`
	Status          string     `json:"status" form:"status" gorm:"default:available"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	UpcyclerID      uint       `json:"upcycler_id"`
	Upcycler        users.User `json:"upcycler" gorm:"foreignKey:UpcyclerID"`
}
