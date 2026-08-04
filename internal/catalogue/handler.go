package catalogue

import (
	"github.com/gin-gonic/gin"
)

// GetCatalogues godoc
//
//	@Summary		Get all catalogues
//	@Description	Retrieve all catalogue entries
//	@Tags			Catalogue
//	@Produce		json
//	@Success		200	{array}	catalogue.Catalogue
//	@Router			/catalogue [get]
func GetCatalogues(c *gin.Context) {
	catalogue := GetAllCatalogues()
	c.JSON(200, catalogue)
}

// CreateCatalogue godoc
//
//	@Summary		Create catalogue
//	@Description	Create a new catalogue entry
//	@Tags			Catalogue
//	@Accept			json
//	@Produce		json
//	@Param			request	body		catalogue.Catalogue	true	"Catalogue data"
//	@Success		201		{object}	catalogue.Catalogue
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/catalogue [post]
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

// UpdateCatalogue godoc
//
//	@Summary		Update catalogue
//	@Description	Update a catalogue entry
//	@Tags			Catalogue
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Catalogue ID"
//	@Param			request	body		catalogue.Catalogue	true	"Updated catalogue"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/catalogue/{id} [patch]
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

// DeleteCatalogue godoc
//
//	@Summary		Delete catalogue
//	@Description	Delete a catalogue entry
//	@Tags			Catalogue
//	@Produce		json
//	@Param			id	path	int	true	"Catalogue ID"
//	@Success		200	{object}	map[string]string
//	@Failure		400	{object}	map[string]string
//	@Router			/catalogue/{id} [delete]
func DeleteCatalogue(c *gin.Context) {
	id := c.Param("id")

	err := DeleteCatalogueService(id)

	if err != nil {
		c.JSON(400, gin.H{
			"eror": "Catalogue not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Catalogue deleted successfully",
	})
}
