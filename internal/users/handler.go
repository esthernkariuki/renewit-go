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
