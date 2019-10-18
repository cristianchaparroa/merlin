package api

import (
	"fmt"
	"net/http"

	"github.com/cristianchaparroa/merlin/bycles/models"
	"github.com/cristianchaparroa/merlin/bycles/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type BicycleHandler struct {
	db *gorm.DB
}

func NewBicycleHandler(db *gorm.DB) *BicycleHandler {
	return &BicycleHandler{db: db}
}

// GetAll
func (h BicycleHandler) GetAll(c *gin.Context) {
	r := repositories.NewBicycleRepository(h.db)
	bs := r.FindAll()
	c.JSON(200, bs)
}

func (h BicycleHandler) GetById(c *gin.Context) {

}

// Create
func (h BicycleHandler) Create(c *gin.Context) {

	var req models.Bicycle

	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusMethodNotAllowed, fmt.Sprintf("Invalid input%v", err.Error()))
		return
	}

	r := repositories.NewBicycleRepository(h.db)
	nm, err := r.Create(&req)

	if err != nil {

	}

	c.JSON(201, nm)
}

// Update
func (h BicycleHandler) Update(c *gin.Context) {

}
