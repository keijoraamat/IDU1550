package DAL

import models "github.com/keijoraamat/IDU1550/ylesanne5/models"

type ItemRepositoryInterface interface {
	GetAllItems() []models.Item
	CreateItem(item *models.Item) uint
}
