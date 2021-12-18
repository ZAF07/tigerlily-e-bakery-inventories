package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tigerlily-e-bakery-server/internal/pkg/env"
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
	
		db, err := gorm.Open("postgres",  env.GetDBEnv())
	if err != nil {
		log.Fatalf("Error connectiong to Database : %+v", err)
	}
	fmt.Println("SUCCEEDED CONNECTING TO DB")

	ORM = db
}