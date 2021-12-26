package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tigerlily-e-bakery-server/internal/db"
	"github.com/tigerlily-e-bakery-server/internal/pkg/logger"
	rpc "github.com/tigerlily-e-bakery-server/internal/pkg/protos"
	"github.com/tigerlily-e-bakery-server/internal/service/checkout"
)

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
		log.Fatalf("error binding req struct : %+v", err)
	}
	fmt.Printf("HERE : %+v", &req)
	ctx := context.Background()
	service := checkout.NewCheckoutService(a.db)

	resp, err := service.Checkout(ctx, &req)
	if err != nil {
		a.logs.ErrorLogger.Println("[CONTROLLER] Error getting response")
	}

	c.JSON(http.StatusOK,
	gin.H{
		"message": "Success checkout",
		"status": http.StatusOK,
		"data": resp,
	})
}