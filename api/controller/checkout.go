package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CheckoutAPI struct {}

func (a CheckoutAPI) Checkout(c *gin.Context) {
	c.JSON(http.StatusOK,
	gin.H{
		"message": "Success",
		"status": http.StatusOK,
		"data": "Checkout page",
	})
}