package auth

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required"`
}

type LoginRequest struct {
	Phone    string `json:"Phone" binding:"required"`
	Password string `json:"Password" binding:"required"`
}
