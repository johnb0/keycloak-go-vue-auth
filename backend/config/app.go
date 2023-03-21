package config

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Keycloak Keycloak
	Server   Server
}

type Keycloak struct {
	Url           string `validate:"required"`
	Realm         string `validate:"required"`
	AdminUser     string `validate:"required"`
	AdminPassword string `validate:"required"`
	ClientName    string `validate:"required"`
}

type Server struct {
	Host string
	Port int `validate:"required"`
}

func GetAppConfig(additionalDirectories ...string) (*AppConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	for _, d := range additionalDirectories {
		viper.AddConfigPath(d)
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("reading in config: %w", err)
	}

	conf := &AppConfig{}
	if err := viper.Unmarshal(conf); err != nil {
		return nil, fmt.Errorf("unmarshaling config into struct: %w", err)
	}

	if err := validator.New().Struct(conf); err != nil {
		return nil, fmt.Errorf("validating config: %w", err)
	}

	return conf, nil
}
