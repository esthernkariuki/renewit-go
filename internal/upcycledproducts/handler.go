package upcycledproducts

import (
	"github.com/gin-gonic/gin"
)

func GetUpcycledProducts(c *gin.Context) {
	upcycledProducts := GetAllUpcycledProducts()
	c.JSON(200, upcycledProducts)
}

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
