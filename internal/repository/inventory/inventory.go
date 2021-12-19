package inventory

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tigerlily-e-bakery-server/internal/db/models"
)

// Create an interface to prevent unwanted use of these methods

type InventoryRepo struct {
	db *gorm.DB
}

// Receives the db instance as argument and sets it in the struct before returning the struct itself
func NewInventoryRepo(db *gorm.DB) *InventoryRepo {
	return &InventoryRepo{
		db: db,
	}
}

func (m InventoryRepo) GetAllInventories(limit, offset int32) (items []*models.Sku, err error) {
	// m.db.Begin()
	m.db.Debug().Find(&items)
	// defer m.db.Close()
	return
}