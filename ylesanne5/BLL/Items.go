package BLL

import (
	models "github.com/keijoraamat/IDU1550/ylesanne5/models"

	"github.com/keijoraamat/IDU1550/ylesanne5/DAL"
)

// Service Layer
type ItemService struct {
	Repo *DAL.ItemRepository
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
