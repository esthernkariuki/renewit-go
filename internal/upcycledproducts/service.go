package upcycledproducts

import "errors"

func GetAllUpcycledProducts() []Upcycledproduct {
	return FetchUpcyclerProducts()
}

func CreateUpcycledProductsService(upcycledproduct *Upcycledproduct) error {
	return SaveUpcycledProduct(upcycledproduct)
}
func UpdateUpcycledProductsService(
	id string,
	userID uint,
	upcycledproduct *Upcycledproduct,
) error {

	var existing Upcycledproduct

	err := GetUpcycledProductByID(id, &existing)

	if err != nil {
		return err
	}

	if existing.UpcyclerID != userID {
		return errors.New("you do not own this product")
	}

	return UpdateUpcycledProductRepository(id, upcycledproduct)
}

func DeleteUpcycledProductsService(
	id string,
	userID uint,
) error {

	var product Upcycledproduct

	err := GetUpcycledProductByID(id, &product)

	if err != nil {
		return err
	}

	if product.UpcyclerID != userID {
		return errors.New("you do not own this product")
	}

	return DeleteUpcycledProductsRepository(id)
}
