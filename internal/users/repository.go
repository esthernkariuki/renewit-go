package users

import "renewit-go/database"

func FetchUsers() []User {
	var users []User
	database.DB.Find(&users)
	return users
}
func SaveUser(user *User) error {
	result := database.DB.Create(user)
	return result.Error
}
