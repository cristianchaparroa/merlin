package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Bicycle is the Bicycle representation in the database
type Bicycle struct {
	ID    uuid.UUID `gorm:"type:uuid;primary_key;"`
	Model string    `json:"model"`
	Frame string    `json:"frame"`
	Color string    `json:"color"`
	Price int64     `json:"price"`
	Rin   int       `json:"rin"`

	CreatedAt time.Time `json:"createdAt"  time_format:"2006-01-02 15:04:05"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (s *Bicycle) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid.String())
}
