package PL

import "github.com/keijoraamat/IDU1550/ylesanne5/models"

type ItemService interface {
	GetAllItems() []models.Item
	CreateItem(item *models.Item) uint
}