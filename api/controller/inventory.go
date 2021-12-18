package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tigerlily-e-bakery-server/internal/service"
)

type InventoryAPI struct {}

func (a InventoryAPI) GetAllInventories(c *gin.Context) {

	service := service.NewInventoryService()
	fmt.Printf("service %+v", service)
	ctx := context.Background()
	resp, err := service.GetAllInventories(ctx, "g")
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



