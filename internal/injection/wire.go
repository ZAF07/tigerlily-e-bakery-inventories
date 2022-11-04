//go:build wireinject
// +build wireinject

package injection

import (
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/config"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/injection/file"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func LoadGeneralConfig() *config.GeneralConfig {
	wire.Build(file.GeneralConfigProvider)
	return &config.GeneralConfig{}
}

func GetPostgresCredentials() string {
	wire.Build(file.PostgresConnectionStringProvider)
	return LoadGeneralConfig().PostgresDBCredentials.GetPostgresDBString()
}

func GetPostgresDBInstance() *gorm.DB {
	return LoadGeneralConfig().PostgresDB
}
