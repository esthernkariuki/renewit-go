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

	err := UpdateMaterialService(id, &material)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "Material not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Material updated successfully",
	})
}

func DeleteMaterial(c *gin.Context) {
	id := c.Param("id")

	err := DeleteMaterialService(id)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "Material not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Material deleted successfully",
	})

}
