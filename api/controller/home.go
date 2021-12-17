package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BakeryAPI struct {}

func (a BakeryAPI) Index(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"message": "Success",
			"status": http.StatusOK,
			"data": "This is all the pastries",
		},
	)
}

func (a BakeryAPI) GetAllPastries(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"message": "Success",
			"status": http.StatusOK,
			"data": "This is all the pastries",
		},
	)
}


