package db

import (
	"log"

	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/pkg/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var ORM *gorm.DB

type Db struct {
	db *gorm.DB
}

func NewDB(conn string) *gorm.DB {
	connectDB(conn)
	return ORM
}

func connectDB(conn string) {

	logs := logger.NewLogger()
	db, err := gorm.Open("postgres", conn)
	// ❌ NOT IN USE CURRENTLY. USING VIPER AND CONFIG FILE
	// db, err := gorm.Open("postgres",  env.GetDBEnv())
	if err != nil {
		logs.ErrorLogger.Printf("Couldn't connect to Database %+v", err)
		log.Fatalf("Error connectiong to Database : %+v", err)
	}
	logs.InfoLogger.Println("Successfully connected to Database")

	ORM = db
}
