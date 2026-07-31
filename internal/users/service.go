package users

func GetAllUsers() []User {

	return FetchUsers()
}
func CreateUserService(user *User) error {
	return SaveUser(user)
}

func UpdateUsersService(id string, user *User) error {
	return updateUsersRepository(id, user)
}

func DeleteUserService(id string) error {
	return DeleteUsersRepository(id)
}
