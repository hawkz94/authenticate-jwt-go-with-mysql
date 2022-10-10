package migrations

import (
	"tecmentor-api/models"

	"gorm.io/gorm"
)

func RunMigrate(db *gorm.DB) {
	db.AutoMigrate(models.Companies{})
	db.AutoMigrate(models.Users{})
}
