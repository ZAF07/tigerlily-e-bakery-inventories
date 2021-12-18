package service

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tigerlily-e-bakery-server/internal/db"
	rpc "github.com/tigerlily-e-bakery-server/protos"
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

func (srv Service) GetAllInventories(ctx context.Context, req *rpc.GetAllInventoriesReq) (resp *rpc.GetAllInventoriesResp, err error) {
	fmt.Println("SERVICE LAYER")

	i := []*rpc.Sku{}

	srv.db.Debug().Find(&i)
	defer srv.db.Close()
	fmt.Printf("DATA :  %+v", i)

	resp = &rpc.GetAllInventoriesResp{
		Inventories: i,
	}

	return
}