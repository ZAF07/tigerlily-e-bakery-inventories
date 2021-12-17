package main

import (
	"log"
	"net"
	"net/http"

	"github.com/tigerlily-e-bakery-server/api/router"

	"github.com/gin-gonic/gin"
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

	s.Serve(l)

}