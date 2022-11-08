package config

import (
	"database/sql"
	"fmt"
)

type AppConfig struct {
	GeneralConfig GeneralConfig `mapstructure:"general_config" json:"general_config"`
}

type GeneralConfig struct {
	Environment      string         `mapstructure:"environment" json:"environment"`
	AppPort          string         `mapstructure:"application_port" json:"application_port"`
	PostgresConfig   PostgresConfig `mapstructure:"postgres_db" json:"postgres_db"`
	PostgresDB       *sql.DB
	CorsAllowOrigins []string `mapstructure:"cors_allow_origin" json:"cors_allow_origin"`
}

type PostgresConfig struct {
	Environment string `mapstructure:"environment" json:"environment"`
	Port        string `mapstructure:"port" json:"port"`
	Host        string `mapstructure:"host" json:"host"`
	Name        string `mapstructure:"name" json:"name"`
	SSL         string `mapstructure:"ssl" json:"ssl"`
	User        string `mapstructure:"user" json:"user"`
	Password    string `mapstructure:"password" json:"password"`
	MaxConns    int    `mapstructure:"max_conns" json:"max_conns"`
}

func (pdb *PostgresConfig) GetPostgresDBString() string {

	if pdb.Password != "" {
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", pdb.Host, pdb.User, pdb.Password, pdb.Name, pdb.Port, pdb.SSL)
	}
	return fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s", pdb.Host, pdb.User, pdb.Name, pdb.Port, pdb.SSL)

}

func (g *GeneralConfig) GetPostgresDBInstance() *sql.DB {
	return g.PostgresDB
}

func InitConfig() *AppConfig {
	return loadConfig()
}
