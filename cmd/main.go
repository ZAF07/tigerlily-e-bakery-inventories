package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/router"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/api/rpc"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/cmd/app"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/config"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/injection"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/pkg/logger"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/service/inventory"
	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

func main() {
	logs := logger.NewLogger()
	logs.InfoLogger.Println("Starting up server ...")

	// ❌ CURRENTLY NOT IN USE, USING VIPER TO GET CONFIG VALUES FROM YAML FILE
	// Set ENV vars
	// env.SetEnv()

	conf := config.InitConfig()
	app.InitAppDependencies(conf)

	port := fmt.Sprintf(":%s", conf.GeneralConfig.AppPort)

	// Spin up the main server instance
	// var port = flag.String("port", ":8000", "Port to listen on")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logs.ErrorLogger.Println("Something went wrong in the server startup")
		log.Fatalf("Error connecting tcp port %s", port)
	}
	logs.InfoLogger.Println("Successfull server init")

	// Start a new multiplexer passing in the main server
	m := cmux.New(lis)

	// Listen for HTTP requests first
	// If request headers don't specify HTTP, next mux would handle the request
	httpListener := m.Match(cmux.HTTP1Fast())
	grpclistener := m.Match(cmux.Any())

	// Run GO routine to run both servers at diff processes at the same time
	go serveGRPC(grpclistener)
	go serveHTTP(httpListener)

	fmt.Printf("Inventory Service Running@%v\n", lis.Addr())

	if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
		log.Fatalf("MUX ERR : %+v", err)
	}

}

// GRPC Server initialisation
func serveGRPC(l net.Listener) {
	// Create a new GRPC server instance
	grpcServer := grpc.NewServer()

	// Register GRPC stubs (pass the GRPC server and the initialisation of the service layer)
	// 💡 NOTE: To Register a GRPC service, we pass in the GRPC server and our client struct. This client struct is the service layer
	rpc.RegisterInventoryServiceServer(grpcServer, inventory.NewInventoryService(injection.GetPostgresDBInstance()))
	// rpc.RegisterInventoryServiceServer(grpcServer, inventory.NewInventoryService(db.NewDB(injection.GetPostgresCredentials())))

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("error running GRPC server %+v", err)
	}
}

// HTTP Server initialisation (using gin)
func serveHTTP(l net.Listener) {
	h := gin.Default()
	router.Router(h)
	s := &http.Server{
		Handler: h,
	}
	if err := s.Serve(l); err != cmux.ErrListenerClosed {
		log.Fatalf("error serving HTTP : %+v", err)
	}
}
