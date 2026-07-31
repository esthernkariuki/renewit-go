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
func updateUsersRepository(id string, users *User) error {
	var existing User

	result := database.DB.First(&existing, id)

	if result.Error != nil {
		return result.Error
	}

	existing.Name = users.Name
	existing.Phone = users.Phone
	existing.Role = users.Role

	result = database.DB.Save(&existing)
	return result.Error

}

func DeleteUsersRepository(id string) error {
	var user User

	result := database.DB.First(&user, id)

	if result.Error != nil {
		return result.Error
	}

	result = database.DB.Delete(user)

	return result.Error

}
