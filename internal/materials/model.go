package materials

import (
	"renewit-go/internal/users"
	"time"
)

type Material struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Type      string     `json:"type" binding:"required"`
	Quantity  int        `json:"quantity" binding:"required,min=1"`
	Condition string     `json:"condition" binding:"required"`
	Image     string     `json:"image"`
	TraderID  uint       `json:"trader_id"`
	Trader    users.User `json:"-" gorm:"foreignKey:TraderID"`
	ListedAt  time.Time  `json:"listed_at"`
}
