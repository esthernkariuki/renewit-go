package upcycledproducts

import (
	"github.com/gin-gonic/gin"
)

// GetUpcycledProducts godoc
//
//	@Summary		Get all upcycled products
//	@Description	Retrieve all upcycled products
//	@Tags			Upcycled Products
//	@Produce		json
//	@Success		200	{array}	upcycledproducts.Upcycledproduct
//	@Router			/upcycled-products [get]
func GetUpcycledProducts(c *gin.Context) {
	upcycledProducts := GetAllUpcycledProducts()
	c.JSON(200, upcycledProducts)
}

// CreateUpcycledProducts godoc
//
//	@Summary		Create upcycled product
//	@Description	Create a new upcycled product (Upcycler only)
//	@Tags			Upcycled Products
//	@Accept			json
//	@Produce		json
//	@Param			request	body		upcycledproducts.Upcycledproduct	true	"Upcycled product data"
//	@Success		201		{object}	upcycledproducts.Upcycledproduct
//	@Failure		400		{object}	map[string]string
//	@Failure		401		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Security		BearerAuth
//	@Router			/upcycled-products [post]
func CreateUpcycledProducts(c *gin.Context) {
	var upcycledProduct Upcycledproduct

	if err := c.ShouldBindJSON(&upcycledProduct); err != nil {
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

	upcycledProduct.UpcyclerID = uint(userID.(float64))

	err := CreateUpcycledProductsService(&upcycledProduct)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create upcycled product",
		})
		return
	}
	c.JSON(201, upcycledProduct)

}

// UpdateUpcycledProducts godoc
//
//	@Summary		Update upcycled product
//	@Description	Update your own upcycled product
//	@Tags			Upcycled Products
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int								true	"Product ID"
//	@Param			request	body		upcycledproducts.Upcycledproduct	true	"Updated product"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{object}	map[string]string
//	@Failure		401		{object}	map[string]string
//	@Failure		403		{object}	map[string]string
//	@Security		BearerAuth
//	@Router			/upcycled-products/{id} [patch]
func UpdateUpcycledProducts(c *gin.Context) {
	id := c.Param("id")

	var upcycledProduct Upcycledproduct

	if err := c.ShouldBindJSON(&upcycledProduct); err != nil {
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

	err := UpdateUpcycledProductsService(id, uint(userID.(float64)), &upcycledProduct)

	if err != nil {
		c.JSON(403, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Upcycled product updated successfully",
	})
}

// DeleteUpcycledProducts godoc
//
//	@Summary		Delete upcycled product
//	@Description	Delete your own upcycled product
//	@Tags			Upcycled Products
//	@Produce		json
//	@Param			id	path	int	true	"Product ID"
//	@Success		200	{object}	map[string]string
//	@Failure		401	{object}	map[string]string
//	@Failure		403	{object}	map[string]string
//	@Security		BearerAuth
//	@Router			/upcycled-products/{id} [delete]
func DeleteUpcycledProducts(c *gin.Context) {

	id := c.Param("id")

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	err := DeleteUpcycledProductsService(id, uint(userID.(float64)))

	if err != nil {
		c.JSON(403, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Upcycled product deleted successfully",
	})
}
