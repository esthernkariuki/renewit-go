package catalogue

import "renewit-go/database"

func FetchCatalogue() []Catalogue {
	var catalogues []Catalogue
	database.DB.Find(&catalogues)
	return catalogues
}

func SaveCatalogue(catalogue *Catalogue) error {
	result := database.DB.Create(catalogue)
	return result.Error

}

func UpdateCatalogueRepository(id string, catalogue *Catalogue) error {
	var existing Catalogue
	result := database.DB.First(&catalogue, id)

	if result.Error != nil {
		return result.Error
	}
	existing.MaterialType = catalogue.MaterialType
	existing.PricePerKg = catalogue.PricePerKg
	existing.LastUpdateDate = catalogue.LastUpdateDate
	existing.MaterialDescription = catalogue.MaterialDescription

	result = database.DB.Save(&existing)

	return result.Error
}

func DeleteCatalogueRepository(id string) error {
	var catalogue Catalogue

	result := database.DB.First(&catalogue, id)

	if result.Error != nil {
		return result.Error
	}
	result = database.DB.Delete(&catalogue)

	return result.Error
}
