package auth

import (
	"renewit-go/database"
	"renewit-go/internal/users"
)

func FindUserByPhone(phone string) (*users.User, error) {

	var user users.User

	result := database.DB.
		Where("phone = ?", phone).
		First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func CreateUser(user *users.User) error {

	result := database.DB.Create(user)

	return result.Error
}
