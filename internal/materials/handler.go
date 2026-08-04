package materials

import "github.com/gin-gonic/gin"

// GetMaterials godoc
//
//	@Summary		Get all materials
//	@Description	Retrieve all listed materials
//	@Tags			Materials
//	@Produce		json
//	@Success		200	{array}	materials.Material
//	@Router			/materials [get]
func GetMaterials(c *gin.Context) {
	materials := GetAllMaterials()
	c.JSON(200, materials)
}

// CreateMaterials godoc
//
//	@Summary		Create material
//	@Description	Create a new material (Trader only)
//	@Tags			Materials
//	@Accept			json
//	@Produce		json
//	@Param			request	body		materials.Material	true	"Material data"
//	@Success		201		{object}	materials.Material
//	@Failure		400		{object}	map[string]string
//	@Failure		401		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Security		BearerAuth
//	@Router			/materials [post]
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

// UpdateMaterial godoc
//
//	@Summary		Update material
//	@Description	Update your own material
//	@Tags			Materials
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Material ID"
//	@Param			request	body		materials.Material	true	"Updated material"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		401		{object}	map[string]string
//	@Failure		403		{object}	map[string]string
//	@Security		BearerAuth
//	@Router			/materials/{id} [patch]
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

// DeleteMaterial godoc
//
//	@Summary		Delete material
//	@Description	Delete your own material
//	@Tags			Materials
//	@Produce		json
//	@Param			id	path	int	true	"Material ID"
//	@Success		200	{object}	map[string]string
//	@Failure		401	{object}	map[string]string
//	@Failure		403	{object}	map[string]string
//	@Security		BearerAuth
//	@Router			/materials/{id} [delete]
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
