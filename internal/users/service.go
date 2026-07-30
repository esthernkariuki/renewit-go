package users

func GetAllUsers() []User {

	return FetchUsers()
}
func CreateUserService(user *User) error {
	return SaveUser(user)
}
