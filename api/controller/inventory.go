package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tigerlily-e-bakery-server/internal/service"
	rpc "github.com/tigerlily-e-bakery-server/protos"
)

type InventoryAPI struct {}

func (a InventoryAPI) GetAllInventories(c *gin.Context) {

	limit , lErr := strconv.Atoi(c.Query("limit"))
	if lErr != nil {
		fmt.Printf("Error converting limit to int : %+v", lErr)
	}
		offset , oErr := strconv.Atoi(c.Query("offset"))
	if oErr != nil {
		fmt.Printf("Error converting offset to int : %+v", oErr)
	}
	
	req := rpc.GetAllInventoriesReq{
		Limit: int32(limit),
		Offset: int32(offset),
	}

	service := service.NewInventoryService()
	fmt.Printf("service %+v", service)
	ctx := context.Background()
	resp, err := service.GetAllInventories(ctx, &req)
	if err != nil {
		log.Fatalf("Error getting response from service layer")
	}

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



