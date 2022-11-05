package providers

import (
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/config"
	"github.com/jinzhu/gorm"
)

func ApplicationConfigProvider() *config.AppConfig {
	return config.InitConfig()
}

func GeneralConfigProvider() *config.GeneralConfig {
	return &config.InitConfig().GeneralConfig
}

func PostgresConnectionStringProvider() string {
	return config.InitConfig().GeneralConfig.PostgresDBCredentials.GetPostgresDBString()
}

func PostgresDBProvider() *gorm.DB {
	return config.InitConfig().GetPostgresDBInstance()
}
