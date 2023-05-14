package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ridhohrnf/REST-API-Upload-Photo-With-Auth/helpers"
	"github.com/ridhohrnf/REST-API-Upload-Photo-With-Auth/models"
)

type PhotoController struct {
	db *gorm.DB
}

func NewPhotoController(db *gorm.DB) *PhotoController {
	return &PhotoController{db}
}

func (pc *PhotoController) CreatePhoto(c *gin.Context) {
	// Handle photo creation logic
}

func (pc *PhotoController)
