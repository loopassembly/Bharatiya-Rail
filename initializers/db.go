package initializers

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"railway-bac/models"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("/data/railway.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate all the models
	err = DB.AutoMigrate(
		&models.User{},
		&models.CBCMaterialRecord{},
	)
	if err != nil {
		log.Fatal("Failed to auto-migrate models:", err)
	}
}
