package providers

import (
	"database/sql"

	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/config"
)

func ApplicationConfigProvider() *config.AppConfig {
	return config.InitConfig()
}

func GeneralConfigProvider() *config.GeneralConfig {
	return &config.InitConfig().GeneralConfig
}

func PostgresConnectionStringProvider() string {
	return config.InitConfig().GeneralConfig.PostgresConfig.GetPostgresDBString()
}

func PostgresDBProvider() *sql.DB {
	return config.InitConfig().GeneralConfig.GetPostgresDBInstance()
}
