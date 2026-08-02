package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"renewit-go/internal/users"
)

func Register(request RegisterRequest) error {

	// Check if a user with this phone already exists
	existingUser, err := FindUserByPhone(request.Phone)

	// User already exists
	if err == nil && existingUser != nil {
		return errors.New("phone number already registered")
	}

	// Something went wrong with the database
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	// Create the database model
	user := users.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: string(hashedPassword),
		Role:     request.Role,
	}

	// Save the user
	err = CreateUser(&user)
	if err != nil {
		return err
	}

	return nil
}

func Login(request LoginRequest) (string, error) {

	user, err := FindUserByPhone(request.Phone)
	if err != nil {
		return "", errors.New("invalid phone or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(request.Password),
	)

	if err != nil {
		return "", errors.New("invalid phone or password")
	}

	token, err := GenerateJWT(
		user.ID,
		user.Phone,
		user.Role,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}
