package users

import (
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := GetAllUsers()
	c.JSON(200, users)
}
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
