package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/db"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/pkg/logger"
	rpc "github.com/ZAF07/tigerlily-e-bakery-inventories/internal/pkg/protos"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/service/checkout"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// hello

type CheckoutAPI struct {
	db *gorm.DB
	logs logger.Logger
}

// Init the DB here (open a connection to the DB) and pass it along to service and repo layer
func NewCheckoutAPI() *CheckoutAPI {
	return &CheckoutAPI{
		db: db.NewDB(),
		logs: *logger.NewLogger(),
	}
}

func (a CheckoutAPI) Checkout(c *gin.Context) {
	a.logs.InfoLogger.Println("[CONTROLLER] Checkout API running")

	var req rpc.CheckoutReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Printf("error binding req struct : %+v", err)
	}
	fmt.Printf("HERE : %+v", req.CheckoutItems[0])
	ctx := context.Background()
	service := checkout.NewCheckoutService(a.db)

	// PROPERLY HANDLE ERROR FOR WHEN DB ERROR 
	resp, err := service.Checkout(ctx, &req)
	if err != nil {
		a.logs.ErrorLogger.Println("[CONTROLLER] Error getting response")
		a.logs.InfoLogger.Printf("[CONTROLLER] Status of resp value: %+v\n",resp)
		c.JSON(http.StatusInternalServerError,
		gin.H{
		"message": "Error checkout",
		"status": http.StatusInternalServerError,
		"data": resp,
	})
	return
	}

	c.JSON(http.StatusOK,
	gin.H{
		"message": "Success checkout",
		"status": http.StatusOK,
		"data": resp,
	})
}