package db

import (
	"log"

	"database/sql"

	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/injection"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/pkg/logger"
	_ "github.com/lib/pq"

	// "gorm.io/driver/postgres"
	// _ "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// 💡 NEEDS REFACTORING HERE. WORKS BUT A LITTLE MESSY HOW I AM INITIALISING THE DATABASE

var ORM *gorm.DB
var PgDB *sql.DB

type Db struct {
	db *gorm.DB
}

func NewDB(conn string) *gorm.DB {
	connectDB(conn)
	return ORM
}

func connectDB(conn string) {

	logs := logger.NewLogger()
	// db, err := gorm.Open(postgres.Open(conn))

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

func NewPgDatabase(conn string) *sql.DB {
	logs := logger.NewLogger()

	config := injection.LoadGeneralConfig()
	maxConns := config.PostgresConfig.MaxConns

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalf("ERROR CONNECTION TO DATABASE: %+v\n", err)
	}

	if pingErr := db.Ping(); err != nil {
		log.Fatalf("Ping database error: %+v\n", pingErr)
	}

	db.SetMaxOpenConns(maxConns)

	logs.InfoLogger.Println("Successfully connected to database 👍")

	PgDB = db
	return db
}
