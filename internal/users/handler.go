package users

import (
	"github.com/gin-gonic/gin"
)

// GetUsers godoc
//
//	@Summary		Get all users
//	@Description	Retrieve all registered users
//	@Tags			Users
//	@Produce		json
//	@Success		200	{array}		users.User
//	@Router			/users [get]
func GetUsers(c *gin.Context) {
	users := GetAllUsers()
	c.JSON(200, users)
}

// CreateUser godoc
//
//	@Summary		Create a user
//	@Description	Create a new user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		users.User	true	"User data"
//	@Success		201		{object}	users.User
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/users [post]
func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := CreateUserService(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create user",
		})
		return
	}
	c.JSON(201, user)
}

// UpdateUser godoc
//
//	@Summary		Update user
//	@Description	Update an existing user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int			true	"User ID"
//	@Param			request	body		users.User	true	"Updated user"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/users/{id} [patch]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := UpdateUsersService(id, &user)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to update material",
		})
		return
	}
	c.JSON(201, gin.H{
		"message": "User updated successfully",
	})
}

// DeleteUser godoc
//
//	@Summary		Delete user
//	@Description	Delete a user by ID
//	@Tags			Users
//	@Produce		json
//	@Param			id	path	int	true	"User ID"
//	@Success		204	{object}	map[string]string
//	@Failure		409	{object}	map[string]string
//	@Router			/users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := DeleteUserService(id)

	if err != nil {
		c.JSON(409, gin.H{
			"error": "Cannot delete user because they have associated payments.",
		})
		return
	}
	c.JSON(204, gin.H{
		"Message": "User deleted successfully",
	})

}
