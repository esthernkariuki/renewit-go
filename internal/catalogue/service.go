package catalogue

func GetAllCatalogues() []Catalogue {
	return FetchCatalogue()
}

func CreateCatalogueService(catalogue *Catalogue) error {
	return SaveCatalogue(catalogue)
}

func UpdateCatalogueService(id string, catalogue *Catalogue) error {
	return UpdateCatalogueRepository(id, catalogue)
}

func DeleteCatalogueService(id string) error {
	return DeleteCatalogueRepository(id)
}
