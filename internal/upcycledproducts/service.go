package upcycledproducts

func GetAllUpcycledProducts() []Upcycledproduct {
	return FetchUpcyclerProducts()
}

func CreateUpcycledProductsService(upcycledproduct *Upcycledproduct) error {
	return SaveUpcycledProduct(upcycledproduct)
}
func UpdateUpcycledProductsService(id string, upcycledproduct *Upcycledproduct) error {
	return UpdateUpcycledProductRepository(id, upcycledproduct)
}

func DeleteUpcycledProductsService(id string) error {
	return DeleteUpcycledProductsRepository(id)
}
