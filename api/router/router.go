package router

import (
	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/controller"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) *gin.Engine {
	
	// Set CORS config
	r.Use(cors.New(cors.Config{
		AllowCredentials: false,
		// AllowAllOrigins: true,
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTION", "HEAD", "PATCH", "COMMON"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Authorization", "accept", "origin", "Referer", "User-Agent"},
	}))

	r.Use(middleware.CORSMiddleware())

	// Get all inventory
	// Get all invetory by type
	// inventoryAPI := new(controller.InventoryAPI)
	inventoryAPI := controller.NewInventoryAPI()
	inventory := r.Group("inventory")
	{
		inventory.GET("", inventoryAPI.GetAllInventories)
		inventory.GET("/:type",inventoryAPI.GetInventoryByType)


		// Checkout API Endpoint
		checkoutAPI := controller.NewCheckoutAPI()
		checkOut := r.Group("checkout")
		{
			checkOut.POST("", checkoutAPI.Checkout)
		}
	
		// Get user details and past pre-checkout cart item
		cache := r.Group("cache")
		cacheAPI := new(controller.CacheAPI)
		{
			cache.GET("/:user_uuid", cacheAPI.GetUserDetails)
		}
	}

	return r
}