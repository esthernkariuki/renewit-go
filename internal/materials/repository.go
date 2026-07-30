package materials

import "renewit-go/database"

func FetchMaterials() []Material {
	var materials []Material

	database.DB.Preload("Trader").Find(&materials)
	return materials

}

func SaveMaterial(material *Material) error {
	result := database.DB.Create(material)

	return result.Error

}

func UpdateMaterialRepository(id string, material *Material) error {

	var existing Material

	result := database.DB.First(&existing, id)

	if result.Error != nil {
		return result.Error
	}

	existing.Type = material.Type
	existing.Quantity = material.Quantity
	existing.Condition = material.Condition
	existing.Image = material.Image

	result = database.DB.Save(&existing)

	return result.Error
}

func DeleteMaterialRepository(id string) error {

	var material Material

	result := database.DB.First(&material, id)

	if result.Error != nil {
		return result.Error
	}

	result = database.DB.Delete(&material)

	return result.Error
}
