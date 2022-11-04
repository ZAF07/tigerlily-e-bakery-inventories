package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type AppConfig struct {
	GeneralConfig GeneralConfig `mapstructure:"general_config" json:"general_config"`
}

type GeneralConfig struct {
	Environment           string         `mapstructure:"environment" json:"environment"`
	AppPort               string         `mapstructure:"application_port" json:"application_port"`
	PostgresDBCredentials PostgresConfig `mapstructure:"postgres_db" json:"postgres_db"`
	PostgresDB            *gorm.DB
	CorsAllowOrigins      []string `mapstructure:"cors_allow_origin" json:"cors_allow_origin"`
}

type PostgresConfig struct {
	Environment string `mapstructure:"environment" json:"environment"`
	Port        string `mapstructure:"port" json:"port"`
	Host        string `mapstructure:"host" json:"host"`
	Name        string `mapstructure:"name" json:"name"`
	SSL         string `mapstructure:"ssl" json:"ssl"`
	User        string `mapstructure:"user" json:"user"`
	Password    string `mapstructure:"password" json:"password"`
}

func (pdb *PostgresConfig) GetPostgresDBString() string {
	dbString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", pdb.Host, pdb.User, pdb.Password, pdb.Name, pdb.Port, pdb.SSL)

	return dbString
}

func (a *AppConfig) GetPostgresDBInstance() *gorm.DB {
	return a.GeneralConfig.PostgresDB
}

func InitConfig() *AppConfig {
	return loadConfig()
}
