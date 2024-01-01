package BLL

import (
	DAL "github.com/keijoraamat/IDU1550/ylesanne5/DAL"

	models "github.com/keijoraamat/IDU1550/ylesanne5/models"
)

// Service Layer
type ItemService struct {
	Repo DAL.ItemRepositoryInterface
}

func (is *ItemService) GetAllItems() []models.Item {
	return is.Repo.GetAllItems()
}

func (is *ItemService) CreateItem(item *models.Item) uint {
	// do some business logic here
	// or here
	// or here

	var ID = is.Repo.CreateItem(item)
	return ID
}

func (is *ItemService) DeleteItem(id string) {
	is.Repo.DeleteItem(id)
}
