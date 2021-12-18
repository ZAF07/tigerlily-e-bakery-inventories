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
// return &Db{
// 	db: ORM,
// }
}

func connectDB() () {
		dbString := env.GetDBEnv()
		db, err := gorm.Open("postgres",  dbString)
	if err != nil {
		log.Fatalf("Error connectiong to Database : %+v", err)
	}
	fmt.Println("SUCCEEDED CONNECTING TO DB")

	ORM = db
}