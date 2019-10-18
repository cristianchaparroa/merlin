package initializer

import (
	"github.com/cristianchaparroa/merlin/bycles/models"
	"github.com/jinzhu/gorm"
)

func InitModels(db *gorm.DB) {
	db.CreateTable(&models.Bicycle{})
}
