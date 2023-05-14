package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ridhohrnf/REST-API-Upload-Photo-With-Auth/controllers"
	"github.com/ridhohrnf/REST-API-Upload-Photo-With-Auth/middlewares"
)

func SetupUserRoutes(router *gin.Engine, db *gorm.DB) {
	userController := controllers.NewUserController(db)

	router.POST("/users/register", userController.Register)
	router.POST("/users/login", userController.Login)

	authGroup := router.Group("/users")
	authGroup.Use(middlewares.AuthMiddleware())
	{
		authGroup.PUT("/:userId", userController.UpdateUser)
		authGroup.DELETE("/:userId", userController.DeleteUser)
	}
}
