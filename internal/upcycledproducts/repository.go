package upcycledproducts

import (
	"renewit-go/database"
)

func FetchUpcyclerProducts() []Upcycledproduct {
	var upcycledproduct []Upcycledproduct
	database.DB.Preload("Upcycler").Find(&upcycledproduct)

	return upcycledproduct
}

func SaveUpcycledProduct(upcycledproducts *Upcycledproduct) error {
	result := database.DB.Save(&upcycledproducts)
	return result.Error
}
func UpdateUpcycledProductRepository(id string, upcycledproducts *Upcycledproduct) error {
	var existing Upcycledproduct
	result := database.DB.First(&existing, id)

	if result.Error != nil {
		return result.Error
	}

	existing.Image = upcycledproducts.Image
	existing.Quantity = upcycledproducts.Quantity
	existing.Type = upcycledproducts.Type
	existing.Price = upcycledproducts.Price

	result = database.DB.Save(existing)

	return result.Error
}

func DeleteUpcycledProductsRepository(id string) error {

	var upcycledProduct Upcycledproduct

	result := database.DB.First(&upcycledProduct, id)

	if result.Error != nil {
		return result.Error
	}

	result = database.DB.Delete(upcycledProduct)
	return result.Error

}
