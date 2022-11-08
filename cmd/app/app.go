package app

import (
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/config"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/db"
)

func InitAppDependencies(config *config.AppConfig) {

	initPostgres(config.GeneralConfig.PostgresConfig.GetPostgresDBString(), config)
}

func initPostgres(dbConn string, appConfig *config.AppConfig) {
	db := db.NewPgDatabase(dbConn)

	appConfig.GeneralConfig.PostgresDB = db
}
