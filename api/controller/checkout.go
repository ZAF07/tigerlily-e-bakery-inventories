package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tigerlily-e-bakery-server/internal/db"
	"github.com/tigerlily-e-bakery-server/internal/pkg/logger"
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
// c.Set("Access-Control-Allow-Origin", true)	
c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK,
	gin.H{
		"message": "Success",
		"status": http.StatusOK,
		"data": "Checkout page",
	})
}