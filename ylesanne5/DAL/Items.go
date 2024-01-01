package DAL

import (
	"log"

	"github.com/keijoraamat/IDU1550/ylesanne5/models"
	"gorm.io/gorm"
)

// Data Access Layer

type ItemRepository struct {
	DB *gorm.DB
}

func (ir *ItemRepository) GetAllItems() []models.Item {

	var items []models.Item
	result := ir.DB.Find(&items)
	if result.Error != nil {
		log.Println("error getting all items.")
	}
	return items

}

func (ir *ItemRepository) CreateItem(item *models.Item) uint {

	result := ir.DB.Create(&item)
	if result.Error != nil {
		log.Println("error creating item: ", result.Error)
		log.Println("Could not create item: ", &item)
	}
	return item.ID

}

func (ir *ItemRepository) DeleteItem(id string) {

	var item models.Item
	ir.DB.Where("id = ?", id).Delete(&item)

}
