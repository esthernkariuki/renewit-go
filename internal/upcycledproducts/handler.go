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
	err := UpdateUpcycledProductsService(id, &upcycledProduct)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to update upcycled product",
		})
		return
	}
	c.JSON(201, upcycledProduct)
}

func DeleteUpcycledProducts(c *gin.Context) {
	id := c.Param("id")

	err := DeleteUpcycledProductsService(id)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "Upcycled product not found",
		})
		return
	}
	c.JSON(204, gin.H{
		"message": "Upcycled product deleted successfully",
	})
}
