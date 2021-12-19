package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tigerlily-e-bakery-server/api/controller"
)

func Router(r *gin.Engine) *gin.Engine {

	// Set CORS config
	r.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowAllOrigins: true,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Authorization", "accept", "origin", "Referer", "User-Agent"},
	}))

	// Get all inventory
	// Get all invetory by type
	// inventoryAPI := new(controller.InventoryAPI)
	inventoryAPI := controller.NewInventoryAPI()
	inventory := r.Group("inventory")
	{
		inventory.GET("", inventoryAPI.GetAllInventories)
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