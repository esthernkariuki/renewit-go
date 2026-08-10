package upcycledproducts

import (
	"renewit-go/database"
)

func FetchUpcyclerProducts() []Upcycledproduct {
	var upcycledProducts []Upcycledproduct

	database.DB.
		Preload("Upcycler").
		Find(&upcycledProducts)

	return upcycledProducts
}

func SaveUpcycledProduct(product *Upcycledproduct) error {
	result := database.DB.Create(product)
	return result.Error
}

func GetUpcycledProductByID(
	id string,
	product *Upcycledproduct,
) error {
	result := database.DB.First(product, id)
	return result.Error
}

func UpdateUpcycledProductRepository(
	id string,
	product *Upcycledproduct,
) error {

	var existing Upcycledproduct

	result := database.DB.First(&existing, id)

	if result.Error != nil {
		return result.Error
	}

	existing.UpcycledClothes = product.UpcycledClothes
	existing.Description = product.Description
	existing.Image = product.Image
	existing.Quantity = product.Quantity
	existing.Type = product.Type
	existing.Material = product.Material
	existing.Size = product.Size
	existing.Color = product.Color
	existing.Condition = product.Condition
	existing.Location = product.Location
	existing.Price = product.Price
	existing.Status = product.Status

	result = database.DB.Save(&existing)

	return result.Error
}

func DeleteUpcycledProductsRepository(id string) error {

	var product Upcycledproduct

	result := database.DB.First(&product, id)

	if result.Error != nil {
		return result.Error
	}

	result = database.DB.Delete(&product)

	return result.Error
}
