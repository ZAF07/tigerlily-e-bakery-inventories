package app

import (
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/config"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/db"
)

func InitAppDependencies(config *config.AppConfig) {

	initPostgres(config.GeneralConfig.PostgresDBCredentials.GetPostgresDBString(), config)
}

func initPostgres(dbConn string, appConfig *config.AppConfig) {
	db := db.NewDB(dbConn)
	appConfig.GeneralConfig.PostgresDB = db
}
