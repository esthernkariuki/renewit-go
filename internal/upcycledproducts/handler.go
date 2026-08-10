package upcycledproducts

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetUpcycledProducts godoc
//
// @Summary      Get all upcycled products
// @Description  Retrieve all upcycled products
// @Tags         Upcycled Products
// @Produce      json
// @Success      200 {array} Upcycledproduct
// @Router       /upcycled-products [get]
func GetUpcycledProducts(c *gin.Context) {
	upcycledProducts := GetAllUpcycledProducts()

	c.JSON(http.StatusOK, upcycledProducts)
}

// CreateUpcycledProducts godoc
//
// @Summary      Create upcycled product
// @Description  Create a new upcycled product with an optional image
// @Tags         Upcycled Products
// @Accept       multipart/form-data
// @Produce      json
// @Param        upcycled_clothes formData string true "Product name"
// @Param        description formData string false "Product description"
// @Param        image formData file false "Product image"
// @Param        quantity formData int true "Quantity"
// @Param        type formData string true "Product type"
// @Param        material formData string false "Material"
// @Param        size formData string false "Size"
// @Param        color formData string false "Color"
// @Param        condition formData string false "Condition"
// @Param        location formData string false "Location"
// @Param        price formData number true "Price"
// @Param        status formData string false "Product status"
// @Success      201 {object} Upcycledproduct
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Security     BearerAuth
// @Router       /upcycled-products [post]
func CreateUpcycledProducts(c *gin.Context) {

	// =========================================
	// GET USER ID
	// =========================================

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	userIDFloat, ok := userID.(float64)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	// =========================================
	// GET FORM VALUES
	// =========================================

	upcycledClothes := c.PostForm("upcycled_clothes")
	description := c.PostForm("description")
	typeValue := c.PostForm("type")
	material := c.PostForm("material")
	size := c.PostForm("size")
	color := c.PostForm("color")
	condition := c.PostForm("condition")
	location := c.PostForm("location")
	status := c.PostForm("status")

	quantityString := c.PostForm("quantity")
	priceString := c.PostForm("price")

	// =========================================
	// VALIDATE REQUIRED FIELDS
	// =========================================

	if upcycledClothes == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Product name is required",
		})
		return
	}

	if typeValue == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Product type is required",
		})
		return
	}

	if quantityString == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Quantity is required",
		})
		return
	}

	quantity, err := strconv.Atoi(quantityString)

	if err != nil || quantity < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Quantity must be a valid number greater than 0",
		})
		return
	}

	if priceString == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Price is required",
		})
		return
	}

	price, err := strconv.ParseFloat(priceString, 64)

	if err != nil || price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Price must be a valid number",
		})
		return
	}

	// =========================================
	// DEFAULT STATUS
	// =========================================

	if status == "" {
		status = "available"
	}

	// =========================================
	// CREATE PRODUCT
	// =========================================

	upcycledProduct := Upcycledproduct{
		UpcycledClothes: upcycledClothes,
		Description:     description,
		Quantity:        quantity,
		Type:            typeValue,
		Material:        material,
		Size:            size,
		Color:           color,
		Condition:       condition,
		Location:        location,
		Price:           price,
		Status:          status,
		UpcyclerID:      uint(userIDFloat),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// =========================================
	// OPTIONAL IMAGE
	// =========================================

	image, err := c.FormFile("image")

	if err == nil {

		uploadDir := "uploads"

		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create upload directory",
			})
			return
		}

		// Create a safer unique filename
		extension := filepath.Ext(image.Filename)

		filename := strconv.FormatInt(time.Now().UnixNano(), 10) + extension

		filePath := filepath.Join(uploadDir, filename)

		if err := c.SaveUploadedFile(image, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save image",
			})
			return
		}

		// Store URL/path in database
		upcycledProduct.Image = "/uploads/" + filename
	}

	// =========================================
	// SAVE TO DATABASE
	// =========================================

	if err := CreateUpcycledProductsService(&upcycledProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create upcycled product",
		})
		return
	}

	// =========================================
	// RESPONSE
	// =========================================

	c.JSON(http.StatusCreated, upcycledProduct)
}

// UpdateUpcycledProducts godoc
//
// @Summary      Update upcycled product
// @Description  Update your own upcycled product
// @Tags         Upcycled Products
// @Accept       multipart/form-data
// @Produce      json
// @Param        id path int true "Product ID"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      403 {object} map[string]string
// @Security     BearerAuth
// @Router       /upcycled-products/{id} [patch]
func UpdateUpcycledProducts(c *gin.Context) {
	id := c.Param("id")

	var upcycledProduct Upcycledproduct

	if err := c.ShouldBind(&upcycledProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	userIDFloat, ok := userID.(float64)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	err := UpdateUpcycledProductsService(
		id,
		uint(userIDFloat),
		&upcycledProduct,
	)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Upcycled product updated successfully",
	})
}

// DeleteUpcycledProducts godoc
//
// @Summary      Delete upcycled product
// @Description  Delete your own upcycled product
// @Tags         Upcycled Products
// @Produce      json
// @Param        id path int true "Product ID"
// @Success      200 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      403 {object} map[string]string
// @Security     BearerAuth
// @Router       /upcycled-products/{id} [delete]
func DeleteUpcycledProducts(c *gin.Context) {
	id := c.Param("id")

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	userIDFloat, ok := userID.(float64)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	err := DeleteUpcycledProductsService(
		id,
		uint(userIDFloat),
	)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Upcycled product deleted successfully",
	})
}
