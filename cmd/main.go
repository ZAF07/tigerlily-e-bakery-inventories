package main

import (
	"log"
	"net"
	"net/http"

	"github.com/ZAF07/tigerlily-e-bakery-server/api/router"
	"github.com/ZAF07/tigerlily-e-bakery-server/internal/pkg/env"
	"github.com/ZAF07/tigerlily-e-bakery-server/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-server/internal/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	logs := logger.NewLogger()
	logs.InfoLogger.Println("Starting up server ...")
	
	// Set ENV vars
	env.SetEnv()

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		logs.ErrorLogger.Println("Something went wrong in the server startup")
		log.Fatalf("Error connecting tcp port 8000")
	}
	logs.InfoLogger.Println("Successfull server init")

	h := gin.Default()
	h.Use(middleware.CORSMiddleware())
		// Set CORS config
	h.Use(cors.New(cors.Config{
		AllowCredentials: false,
		AllowWildcard: true,
		// AllowAllOrigins: true,
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTION", "HEAD", "PATCH", "COMMON"},
		AllowHeaders: []string{"Content-Type", "Content-Length", "Authorization", "accept", "origin", "Referer", "User-Agent"},
	}))
	router.Router(h)
	s := &http.Server{
		Handler: h,
	}


	s.Serve(l)
}