package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CacheAPI struct {}

func (a CacheAPI) GetUserDetails(c *gin.Context) {
	c.JSON(http.StatusOK,
	gin.H{
		"message" : "Success",
		"status": http.StatusOK,
		"data":"Here are all the user details",
	})
}