package config

type AppConfig struct {
	GeneralConfig GeneralConfig `mapstructure:"general_config" json:"general_config"`
}

type GeneralConfig struct {
	Environment      string         `mapstructure:"environment" json:"environment"`
	PostgresDB       PostgresConfig `mapstructure:"postgres_db" json:"postgres_db"`
	CorsAllowOrigins []string       `mapstructure:"cors_allow_origin" json:"cors_allow_origin"`
}

type PostgresConfig struct {
	Environment string `mapstructure:"environment" json:"environment"`
	Port        string `mapstructure:"port" json:"port"`
	Host        string `mapstructure:"host" json:"host"`
	Name        string `mapstructure:"name" json:"name"`
	SSL         string `mapstructure:"ssl" json:"ssl"`
	User        string `mapstructure:"user" json:"user"`
}

func InitConfig() *AppConfig {
	return loadConfig()
}
