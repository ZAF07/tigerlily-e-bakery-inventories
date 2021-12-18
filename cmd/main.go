package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tigerlily-e-bakery-server/api/router"
	"github.com/tigerlily-e-bakery-server/pkg/env"
)

func main() {

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Error connecting tcp port 8000")
	}

	h := gin.Default()
	router.Router(h)
	s := &http.Server{
		Handler: h,
	}

	// Set ENV vars
	env.SetEnv()

	s.Serve(l)
}