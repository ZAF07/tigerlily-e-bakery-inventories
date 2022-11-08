//go:build wireinject
// +build wireinject

package injection

import (
	"database/sql"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/config"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/injection/providers"
	"github.com/google/wire"
)

func LoadGeneralConfig() *config.GeneralConfig {
	wire.Build(providers.GeneralConfigProvider)
	return &config.GeneralConfig{}
}

func GetPostgresCredentials() string {
	wire.Build(providers.PostgresConnectionStringProvider)
	return LoadGeneralConfig().PostgresConfig.GetPostgresDBString()
}

func GetPostgresDBInstance() *sql.DB {
	wire.Build(providers.PostgresDBProvider)
	return LoadGeneralConfig().GetPostgresDBInstance()
}
