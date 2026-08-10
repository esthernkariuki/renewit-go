package upcycledproducts

import "errors"

func GetAllUpcycledProducts() []Upcycledproduct {
	return FetchUpcyclerProducts()
}

func CreateUpcycledProductsService(
	upcycledProduct *Upcycledproduct,
) error {

	return SaveUpcycledProduct(upcycledProduct)
}

func UpdateUpcycledProductsService(
	id string,
	userID uint,
	upcycledProduct *Upcycledproduct,
) error {

	var existing Upcycledproduct

	err := GetUpcycledProductByID(
		id,
		&existing,
	)

	if err != nil {
		return err
	}

	if existing.UpcyclerID != userID {
		return errors.New(
			"you do not own this product",
		)
	}

	return UpdateUpcycledProductRepository(
		id,
		upcycledProduct,
	)
}

func DeleteUpcycledProductsService(
	id string,
	userID uint,
) error {

	var product Upcycledproduct

	err := GetUpcycledProductByID(
		id,
		&product,
	)

	if err != nil {
		return err
	}

	if product.UpcyclerID != userID {
		return errors.New(
			"you do not own this product",
		)
	}

	return DeleteUpcycledProductsRepository(id)
}
