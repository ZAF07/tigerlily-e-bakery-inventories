package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/db"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/service/inventory"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type InventoryAPI struct {
	db *gorm.DB
	logs logger.Logger
}

// Init the DB here (open a connection to the DB) and pass it along to service and repo layer
func NewInventoryAPI() *InventoryAPI {
	return &InventoryAPI{
		db: db.NewDB(),
		logs: *logger.NewLogger(),
	}
}

func (a InventoryAPI) GetAllInventories(c *gin.Context) {

	a.logs.InfoLogger.Println(" [CONTROLLER] GetAllInventories Request recieved")

	// Serialize and structure incoming data before handing them over to the service layer
	limit , lErr := strconv.Atoi(c.Query("limit"))
	if lErr != nil {
		a.logs.ErrorLogger.Printf("Error converting limit query string to int %+v", lErr)
	}
	offset , oErr := strconv.Atoi(c.Query("offset"))
	if oErr != nil {
		a.logs.ErrorLogger.Printf("Error converting offset query string to int %+v", oErr)
	}

	// Construct the request object 
	req := rpc.GetAllInventoriesReq{
		Limit: int32(limit),
		Offset: int32(offset),
	}

	// Init a new service instance passing the DB instance (service will pass this DB inatance to the repo layer later on)
	service := inventory.NewInventoryService(a.db)
	ctx := context.Background()
	resp, err := service.GetAllInventories(ctx, &req)
	if err != nil {
		log.Fatalf("Error getting response from service layer")
	}

	// Wait for service layer to respond with data before sending them over to the client
	c.JSON(http.StatusOK,
		gin.H{
			"message": "Success",
			"status": http.StatusOK,
			"data": resp,
		},
	)
}

func (a InventoryAPI) GetInventoryByType(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"message": "Success",
			"status": http.StatusOK,
			"data": "This is all the pastries by type",
		},
	)
}



