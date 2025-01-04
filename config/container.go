package config

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Container struct {
	Environments *environments
	Logger       *logrus.Logger
}

type applicationMode string
type ginMode string

const aplicationModeKey applicationMode = "aplicationMode"
const ginModeKey ginMode = "ginMode"

func LoadContainer() (context.Context, *Container, error) {
	// Load Context
	ctx := context.Background()

	// Loads Application Mode Into Context
	ctx = context.WithValue(ctx, aplicationModeKey, applicationMode("release"))

	// Load Logger
	logger := loadLogger(ctx)

	// Load Environment Settings
	environments, err := loadEnv()

	if err != nil {
		return nil, nil, err
	}

	// Load Migrations
	err = LoadMigration(fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", environments.DATABASE_AGENT, environments.DATABASE_USER, environments.DATABASE_PASSWORD, environments.DATABASE_HOST, environments.DATABASE_PORT, environments.DATABASE_SCHEMA, environments.DATABASE_SSLMODE))

	if err != nil {
		return nil, nil, err
	}

	logger.Info("Migrations Loaded")

	// Load Gin Mode Into Context
	ctx = context.WithValue(ctx, ginModeKey, ginMode(environments.GIN_MODE))

	return ctx, &Container{
		Environments: environments,
		Logger:       logger,
	}, nil
}
