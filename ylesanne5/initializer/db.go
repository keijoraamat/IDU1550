package initializer

import (
	models "github.com/keijoraamat/IDU1550/ylesanne5/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDB() (DB *gorm.DB) {
	var err error

	DB, err = gorm.Open(sqlite.Open("3layers.db"))
	if err != nil {
		panic("DB connection failed")
	}

	DB.AutoMigrate(
		&models.Item{},
	)
	return
}
