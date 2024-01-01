package PL

import "github.com/keijoraamat/IDU1550/ylesanne5/models"

type ItemServiceInterface interface {
	GetAllItems() []models.Item
	CreateItem(item *models.Item) uint
	DeleteItem(id string)
}
