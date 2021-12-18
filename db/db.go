package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/tigerlily-e-bakery-server/pkg/env"
)

var ORM *gorm.DB

type Db struct {
	db *gorm.DB
}

func NewDB() (db *Db) {
	connectDB()
return &Db{
	db: ORM,
}
}

// func constructDbString(host, user, password, dbname string, port int) (dbString string) {

// 	// dbString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)
// 	dbString = env.GetDBEnv()

// 	return
// }

func connectDB() () {
		dbString := env.GetDBEnv()
		db, err := gorm.Open("postgres",  dbString)
	if err != nil {
		log.Fatalf("Error connectiong to Database : %+v", err)
	}

	ORM = db
}