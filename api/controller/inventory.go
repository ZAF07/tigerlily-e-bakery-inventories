package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tigerlily-e-bakery-server/internal/service/inventory"
	rpc "github.com/tigerlily-e-bakery-server/protos"

	"github.com/tigerlily-e-bakery-server/internal/db"
)

type InventoryAPI struct {}

func (a InventoryAPI) GetAllInventories(c *gin.Context) {
	// Init the DB here (open a connection to the DB) and pass it along to service and repo layer
	DB := db.NewDB()

	// Serialize and structure incoming data before handing them over to the service layer
	limit , lErr := strconv.Atoi(c.Query("limit"))
	if lErr != nil {
		fmt.Printf("Error converting limit to int : %+v", lErr)
	}
		offset , oErr := strconv.Atoi(c.Query("offset"))
	if oErr != nil {
		fmt.Printf("Error converting offset to int : %+v", oErr)
	}

	// Construct the request object 
	req := rpc.GetAllInventoriesReq{
		Limit: int32(limit),
		Offset: int32(offset),
	}

	// Init a new service instance passing the DB instance (service will pass this DB inatance to the repo layer later on)
	service := inventory.NewInventoryService(DB)
	fmt.Printf("service %+v", service)
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



