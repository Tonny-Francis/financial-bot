package config

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func LoadMigration(databaseURL string) error {
	m, err := migrate.New("file://db/migrations", databaseURL)

	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		return err
	}

	return nil
}

func RollbackMigration(databaseURL string) error {
	m, err := migrate.New("file://db/migrations", databaseURL)

	if err != nil {
		return err
	}

	if err := m.Down(); err != nil {
		return err
	}

	return nil
}