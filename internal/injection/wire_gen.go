// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injection

import (
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/config"
	"github.com/ZAF07/tigerlily-e-bakery-inventories/internal/injection/providers"
	"github.com/jinzhu/gorm"
)

// Injectors from wire.go:

func LoadGeneralConfig() *config.GeneralConfig {
	generalConfig := providers.GeneralConfigProvider()
	return generalConfig
}

func GetPostgresCredentials() string {
	string2 := providers.PostgresConnectionStringProvider()
	return string2
}

// wire.go:

func GetPostgresDBInstance() *gorm.DB {
	return LoadGeneralConfig().PostgresDB
}
