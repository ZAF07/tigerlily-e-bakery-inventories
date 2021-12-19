package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tigerlily-e-bakery-server/api/router"
	"github.com/tigerlily-e-bakery-server/internal/pkg/env"
	"github.com/tigerlily-e-bakery-server/internal/pkg/logger"
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
	router.Router(h)
	s := &http.Server{
		Handler: h,
	}


	s.Serve(l)
}