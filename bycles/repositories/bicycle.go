package repositories

import (
	"github.com/cristianchaparroa/merlin/bycles/models"
	"github.com/jinzhu/gorm"
)

type IBicycleRepository interface {
	Create(b *models.Bicycle) (models.Bicycle, error)
	Delete(id string)
	Update(b models.Bicycle) models.Bicycle
	FindByID() *models.Bicycle
	FindAll() []models.Bicycle
}

type BicycleRepository struct {
	db *gorm.DB
}

func NewBicycleRepository(db *gorm.DB) *BicycleRepository {
	return &BicycleRepository{db: db}
}

func (b *BicycleRepository) Create(m *models.Bicycle) (*models.Bicycle, error) {
	b.db.Create(m)
	return m, nil
}

func (b *BicycleRepository) Delete(m models.Bicycle) (*models.Bicycle, error) {
	return nil, nil
}

func (b *BicycleRepository) Update(m models.Bicycle) (*models.Bicycle, error) {
	return nil, nil
}

func (b *BicycleRepository) FindByID(id string) *models.Bicycle {

	var s models.Bicycle
	b.db.Where("id = ?", id).First(&s)
	return &s
}

func (b *BicycleRepository) FindAll() []models.Bicycle {
	ms := make([]models.Bicycle, 0)
	b.db.Find(&ms)
	return ms
}
