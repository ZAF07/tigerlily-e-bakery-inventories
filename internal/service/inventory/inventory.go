package inventory

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/repository/inventory"
	// _ "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

// Data structure representing the service layer. Down below we create methods for this struct to receive
// This struct typically holds the repo layer instance, only made available to use after initialising the repo, passing in the DB instance first
type Service struct {
	db        *sql.DB
	inventory inventory.InventoryRepo
	logs      logger.Logger
	rpc.UnimplementedInventoryServiceServer
}

// This gurantees that Service struct implements the interface
var _ rpc.InventoryServiceServer = (*Service)(nil)

// We initialise a new repo instance at the same time we initialise the service layer
// THE CONTROLLER SHOULD START THE DB INIT AND PASS THE INSTANCE TO SERVICE AND SERVICE TO REPO !!
func NewInventoryService(DB *sql.DB) *Service {
	return &Service{
		db:        DB,
		inventory: *inventory.NewInventoryRepo(DB),
		logs:      *logger.NewLogger(),
	}
}

func (srv Service) GetAllInventories(ctx context.Context, req *rpc.GetAllInventoriesReq) (resp *rpc.GetAllInventoriesResp, err error) {
	srv.logs.InfoLogger.Println(" [SERVICE] GetAllInventories Running service layer")

	// Init a repo instance (the repo should be tied to this service struct)
	in, err := srv.inventory.GetAllInventories(ctx, req.Limit, req.Offset)
	if err != nil {
		srv.logs.ErrorLogger.Printf("Database Error : %+v", err)
		log.Fatalf("Database err %+v", err)
	}

	i := []*rpc.Sku{}
	for _, sku := range in {
		i = append(i, &rpc.Sku{
			Name:        sku.Name,
			Price:       sku.Price,
			SkuId:       sku.SkuID,
			Quantity:    int32(sku.Quantity),
			ImageUrl:    sku.ImageURL,
			Type:        sku.Type,
			Description: sku.Description,
		})
	}

	fmt.Printf("HERE ---> %+v", i[0])
	resp = &rpc.GetAllInventoriesResp{
		Inventories: i,
	}

	return
}
