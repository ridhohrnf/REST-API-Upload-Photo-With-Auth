package database

import (
	"github.com/jinzhu/gorm"
	"github.com/ridhohrnf/REST-API-Upload-Photo-With-Auth/models"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Photo{})
}
