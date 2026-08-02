package materials

import "errors"

func GetAllMaterials() []Material {
	return FetchMaterials()
}

func CreateMaterialService(material *Material) error {
	return SaveMaterial(material)

}

func UpdateMaterialService(id string, userID uint, material *Material) error {

	existing, err := GetMaterialByID(id)
	if err != nil {
		return err
	}

	if existing.TraderID != userID {
		return errors.New("you do not own this material")
	}

	return UpdateMaterialRepository(id, material)
}

func DeleteMaterialService(id string, userID uint) error {

	existing, err := GetMaterialByID(id)
	if err != nil {
		return err
	}

	if existing.TraderID != userID {
		return errors.New("you do not own this material")
	}

	return DeleteMaterialRepository(id)
}
