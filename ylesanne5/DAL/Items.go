package DAL

import (
	"log"

	"github.com/keijoraamat/IDU1550/ylesanne5/models"
	"gorm.io/gorm"
)

// Data Access Layer

type ItemRepository struct {
	DB	*gorm.DB
	items []models.Item
}

func (ir *ItemRepository) GetAllItems() []models.Item {
	// This is where you would retrieve data from a database
	return ir.items
}

func (ir *ItemRepository) CreateItem(item *models.Item) uint  {

	result := ir.DB.Create(&item)
	if result.Error != nil {
		log.Println("Could not create item: ", &item)
	}
	return item.ID
	
}

