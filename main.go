package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ridhohrnf/REST-API-Upload-Photo-With-Auth/app"
	"github.com/ridhohrnf/REST-API-Upload-Photo-With-Auth/routes"
)

func main() {
	router := gin.Default()

	db := app.SetupDatabase()
	defer db.Close()

	routes.SetupUserRoutes(router, db)
	routes.SetupPhotoRoutes(router, db)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
