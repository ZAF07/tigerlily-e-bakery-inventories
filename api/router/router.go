package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tigerlily-e-bakery-server/api/controller"
	"github.com/tigerlily-e-bakery-server/internal/pkg/middleware"
)

func Router(r *gin.Engine) *gin.Engine {

	
	// Set CORS config
	r.Use(cors.New(cors.Config{
		AllowCredentials: false,
		// AllowAllOrigins: true,
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Authorization", "accept", "origin", "Referer", "User-Agent"},
	}))
	
	r.Use(middleware.CORSMiddleware())
	// Get all inventory
	// Get all invetory by type
	// inventoryAPI := new(controller.InventoryAPI)
		checkoutAPI := controller.NewCheckoutAPI()

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
	// checkoutAPI := controller.NewCheckoutAPI()
	checkOut := r.Group("checkout")
	{
		checkOut.GET("/c", checkoutAPI.Checkout)
	}

	return r
}