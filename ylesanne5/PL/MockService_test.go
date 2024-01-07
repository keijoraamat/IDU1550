package PL_test

import (
	"github.com/keijoraamat/IDU1550/ylesanne5/models"
)

type MockService struct{}

func (ms *MockService) GetAllItems() []models.Item {
	return []models.Item{{Name: "item1", Price: 10}, {Name: "item2", Price: 20}}
}

func (ms *MockService) CreateItem(item *models.Item) uint {
	return 1
}
func (ms *MockService) DeleteItem(id string) {}
