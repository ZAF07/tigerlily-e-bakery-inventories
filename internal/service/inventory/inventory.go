package inventory

import (
	"context"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tigerlily-e-bakery-server/internal/repository/inventory"
	rpc "github.com/tigerlily-e-bakery-server/protos"
)

// Data structure representing the service layer. Down below we create methods for this struct to receive
// This struct typically holds the repo layer instance, only made available to use after initialising the repo, passing in the DB instance first
type Service struct {
	db *gorm.DB
	inventory inventory.InventoryRepo
}
// We initialise a new repo instance at the same time we initialise the service layer
// THE CONTROLLER SHOULD START THE DB INIT AND PASS THE INSTANCE TO SERVICE AND SERVICE TO REPO !!
func NewInventoryService(DB *gorm.DB) *Service {
	// n := db.NewDB()

	return&Service{
		db: DB,
		inventory: *inventory.NewInventoryRepo(DB),
	}
}

func (srv Service) GetAllInventories(ctx context.Context, req *rpc.GetAllInventoriesReq) (resp *rpc.GetAllInventoriesResp, err error) {
	fmt.Println("SERVICE LAYER")

	// Init a repo instance (the repo should be tied to this service struct)
	in, err := srv.inventory.GetAllInventories(req.Limit, req.Offset)
	if err != nil {
		log.Fatalf("Database err %+v", err)
	}

	i := []*rpc.Sku{}
	for _, sku := range in {
		i = append(i, &rpc.Sku{
			Name: sku.Name,
			Price: sku.Price,
			SkuId: sku.SkuId,
			ImageUrl: sku.ImageUrl,
			Type: sku.Type,
			Description: sku.Description,
		})
	}

	resp = &rpc.GetAllInventoriesResp{
		Inventories: i,
	}

	return
}