package catalogue

import (
	"github.com/gin-gonic/gin"
)

func GetCatalogues(c *gin.Context) {
	catalogue := GetAllCatalogues()
	c.JSON(200, catalogue)
}
func CreateCatalogue(c *gin.Context) {
	var catalogue Catalogue

	err := c.ShouldBindJSON(&catalogue)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = CreateCatalogueService(&catalogue)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create catalogue",
		})
		return
	}
	c.JSON(201, catalogue)
}

func UpdateCatalogue(c *gin.Context) {
	id := c.Param("id")
	var catalogue Catalogue

	err := c.ShouldBindJSON(&catalogue)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = UpdateCatalogueService(id, &catalogue)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to update catalogue",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Catalogue updated successfully",
	})

}

func DeleteCatalogue(c *gin.Context) {
	id := c.Param("id")

	err := DeleteCatalogueService(id)

	if err != nil {
		c.JSON(400, gin.H{
			"eror": "catalogue not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"error": "Catalogue deleted successfully",
	})
}
