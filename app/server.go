package app

import (
	"github.com/jinzhu/gorm"
	"github.com/ridhohrnf/REST-API-Upload-Photo-With-Auth/database"
)

func SetupDatabase() *gorm.DB {
	db := database.ConnectDB()
	database.RunMigrations(db)
	return db
}
