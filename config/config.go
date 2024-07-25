package config

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost                   string `mapstructure:"DBHOST" validate:"required"`
	DBName                   string `mapstructure:"DBNAME" validate:"required"`
	DBUser                   string `mapstructure:"DBUSER" validate:"required"`
	DBPort                   string `mapstructure:"DBPORT" validate:"required"`
	DBPassword               string `mapstructure:"DBPASSWORD" validate:"required"`
	GrpcPort                 string `mapstructure:"GRPCPORT" validate:"required"`
	GrpcNotificationPort     string `mapstructure:"GrpcNotificationPort" validate:"required"`
	GrpcUserAdminServicePort string `mapstructure:"GrpcUserAdminServicePort" validate:"required"`
}

var envs = []string{
	"DBHOST", "DBNAME", "DBUSER", "DBPORT", "DBPASSWORD", "GRPCPORT", "GrpcNotificationPort", "GrpcUserAdminServicePort",
}

func LoadConfig() (Config, error) {
	var cfg Config
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return cfg, fmt.Errorf("error reading config file: %w", err)
	}

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return cfg, fmt.Errorf("error binding environment variable %s: %w", env, err)
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, fmt.Errorf("error unmarshalling config: %w", err)
	}

	validate := validator.New()
	if err := validate.Struct(&cfg); err != nil {
		return cfg, fmt.Errorf("validation error: %w", err)
	}

	return cfg, nil
}
