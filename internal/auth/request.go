package auth

// package users

// import "time"

// type User struct {
// 	ID        uint      `json:"id"`
// 	Name      string    `json:"name"`
// 	Phone     string    `json:"phone"`
// 	Role      string    `json:"role"`
// 	CreatedAt time.Time `json:"created_at"`
// }

type RegisterRequest struct {
	Name     string `json:"Name" binding:"required"`
	Phone    string `json:"Phone" binding:"required"`
	Password string `json:"Password" binding:"required,min=6"`
	Role     string `json:"Role" binding:"required"`
}

type LoginRequest struct {
	Phone    string `json:"Phone" binding:"required"`
	Password string `json:"Password" binding:"required"`
}
