package service

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tigerlily-e-bakery-server/internal/db"
	"github.com/tigerlily-e-bakery-server/internal/db/models"
)

type Service struct {
	db *gorm.DB
}

func NewInventoryService() *Service {
	n := db.NewDB()

	return&Service{
		db: n,
	}
}


func (srv Service) GetAllInventories(ctx context.Context, req string) (inventory *[]models.Sku, err error) {
	fmt.Println("SERVICE LAYER")

	i := []models.Sku{}

	srv.db.Debug().Find(&i)
	defer srv.db.Close()
	fmt.Printf("DATA :  %+v", i)	
	inventory = &i
	return
}