package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tigerlily-e-bakery-server/api/controller"
)

func Router(r *gin.Engine) *gin.Engine {

	// Get all inventory
	// Get all invetory by type
	inventoryAPI := new(controller.InventoryAPI)
	inventory := r.Group("inventory")
	{
		inventory.GET("/", inventoryAPI.GetAllInventories)
		inventory.GET("/:type",inventoryAPI.GetInventoryByType)
	}

	// Get user details and past pre-checkout cart item
	cache := r.Group("cache")
	cacheAPI := new(controller.CacheAPI)
	{
		cache.GET("/:user_uuid", cacheAPI.GetUserDetails)
	}

	// Checkout API Endpoint
	checkOut := r.Group("checkout")
	checkoutAPI := new(controller.CheckoutAPI)
	{
		checkOut.GET("/", checkoutAPI.Checkout)
	}

	return r
}