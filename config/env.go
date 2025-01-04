package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type environments struct {
	// Config Server HTTP
	GIN_MODE string `validate:"required"`
	PORT     string `validate:"required"`

	// Config Type Env
	GO_ENV string `validate:"required"`

	// Config Database
	DATABASE_AGENT    string `validate:"required"`
	DATABASE_SCHEMA   string `validate:"required"`
	DATABASE_USER     string `validate:"required"`
	DATABASE_PASSWORD string `validate:"required"`
	DATABASE_HOST     string `validate:"required"`
	DATABASE_PORT     string `validate:"required"`
	DATABASE_SSLMODE  string `validate:"required"`
}

func loadEnv() (*environments, error) {
	// Use OS Environment Variables If Config File Is Not Found
	useOS := false

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		useOS = true
	}

	// Load Environment Variables
	env := &environments{
		GIN_MODE: getEnv("GIN_MODE", useOS),
		PORT:     getEnv("PORT", useOS),

		GO_ENV: getEnv("GO_ENV", useOS),

		DATABASE_AGENT:    getEnv("DATABASE_AGENT", useOS),
		DATABASE_SCHEMA:   getEnv("DATABASE_SCHEMA", useOS),
		DATABASE_USER:     getEnv("DATABASE_USER", useOS),
		DATABASE_PASSWORD: getEnv("DATABASE_PASSWORD", useOS),
		DATABASE_HOST:     getEnv("DATABASE_HOST", useOS),
		DATABASE_PORT:     getEnv("DATABASE_PORT", useOS),
		DATABASE_SSLMODE:  getEnv("DATABASE_SSLMODE", useOS),
	}

	// Validate Struct
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(env); err != nil {
		return nil, err
	}

	return env, nil
}

func getEnv(key string, useOS bool) string {
	if useOS {
		return os.Getenv(key)
	}

	return viper.GetString(key)
}
