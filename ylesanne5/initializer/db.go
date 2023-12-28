package initializer

import (
	models "github.com/keijoraamat/IDU1550/ylesanne5/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("3layers.db"))
	if err != nil {
		panic("DB connection failed")
	}
}

func SyncDB() {
	DB.AutoMigrate(
		&models.Item{},
	)
}
