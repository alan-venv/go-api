package configs

import "github.com/spf13/viper"

//? O arquivo de configurações deve ser criado e configurado manualmente

var cfg *config

type config struct {
	Api      apiConfig
	Database dbConfig
}

type apiConfig struct {
	Port string
	Mode string
}

type dbConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	viper.SetDefault("api.port", "8080")
	viper.SetDefault("api.mode", "debug")
	// viper.SetDefault("database.host", "localhost")
	// viper.SetDefault("database.port", "5432")
}

func Load() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic("failed to load config file")
		}
	}

	cfg = new(config)

	cfg.Api = apiConfig{
		Port: viper.GetString("api.port"),
		Mode: viper.GetString("api.mode"),
	}

	cfg.Database = dbConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}
}

func Database() dbConfig {
	return cfg.Database
}

func Api() apiConfig {
	return cfg.Api
}
