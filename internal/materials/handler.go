package materials

import "github.com/gin-gonic/gin"

func GetMaterials(c *gin.Context) {
	materials := GetAllMaterials()
	c.JSON(200, materials)
}

func CreateMaterials(c *gin.Context) {
	var material Material

	if err := c.ShouldBindJSON(&material); err != nil {
		c.JSON(400, gin.H{
			"error":   "Invalid data provided",
			"details": err.Error(),
		})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	material.TraderID = uint(userID.(float64))

	err := CreateMaterialService(&material)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create material",
		})
		return
	}

	c.JSON(201, material)
}

func UpdateMaterial(c *gin.Context) {

	id := c.Param("id")

	var material Material

	if err := c.ShouldBindJSON(&material); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	err := UpdateMaterialService(id, uint(userID.(float64)), &material)

	if err != nil {
		c.JSON(403, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Material updated successfully",
	})
}

func DeleteMaterial(c *gin.Context) {

	id := c.Param("id")

	userID, _ := c.Get("user_id")

	err := DeleteMaterialService(id, uint(userID.(float64)))

	if err != nil {
		c.JSON(403, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Material deleted successfully",
	})
}
