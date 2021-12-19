package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tigerlily-e-bakery-server/internal/pkg/env"
	"github.com/tigerlily-e-bakery-server/internal/pkg/logger"
)

var ORM *gorm.DB

type Db struct {
	db *gorm.DB
}

func NewDB() (*gorm.DB) {
	connectDB()
	return ORM
}

func connectDB() () {

	logs := logger.NewLogger()
		db, err := gorm.Open("postgres",  env.GetDBEnv())
	if err != nil {
		logs.ErrorLogger.Printf("Couldn't connect to Database %+v", err)
		log.Fatalf("Error connectiong to Database : %+v", err)
	}
	logs.InfoLogger.Println("Successfully connected to Database")
	fmt.Println("SUCCEEDED CONNECTING TO DB")

	ORM = db
}