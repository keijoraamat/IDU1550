package BLL_test

import (
	"testing"

	"github.com/keijoraamat/IDU1550/ylesanne5/BLL"
	"github.com/keijoraamat/IDU1550/ylesanne5/models"
)

// MockItemRepository is a mock of ItemRepository
type MockItemRepository struct {
	MockGetAllItems func() []models.Item
	MockCreateItem  func(item *models.Item) uint
	ID              uint
}

func (m *MockItemRepository) GetAllItems() []models.Item {
	return m.MockGetAllItems()
}

func (m *MockItemRepository) CreateItem(item *models.Item) uint {
	item.ID = 1
	return m.MockCreateItem(item)
}

func TestGetAllItems(t *testing.T) {
	mockRepo := new(MockItemRepository)
	mockRepo.MockGetAllItems = func() []models.Item {
		return []models.Item{{Name: "fff", Price: 10}, {Name: "ggg", Price: 2}}
	}
	service := BLL.ItemService{Repo: mockRepo}

	items := service.GetAllItems()

	if len(items) != 2 {
		t.Errorf("GetAllItems was incorrect, got: %d, want: %d.", len(items), 2)
	}
}

func TestCreateItem(t *testing.T) {
	mockRepo := new(MockItemRepository)
	mockRepo.MockCreateItem = func(item *models.Item) uint {
		return 1
	}
	service := BLL.ItemService{Repo: mockRepo}

	item := models.Item{} // Populate this with test data
	id := service.CreateItem(&item)

	if id != 1 {
		t.Errorf("CreateItem was incorrect, got: %d, want: %d.", id, 1)
	}
}
