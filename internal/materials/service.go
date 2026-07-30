package materials

func GetAllMaterials() []Material {
	return FetchMaterials()
}

func CreateMaterialService(material *Material) error {
	return SaveMaterial(material)

}

func UpdateMaterialService(id string, material *Material) error {
	return UpdateMaterialRepository(id, material)
}

func DeleteMaterialService(id string) error {
	return DeleteMaterialRepository(id)
}
