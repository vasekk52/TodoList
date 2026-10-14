package database

import (
	"fmt"

	"todoapp/migrations"

	"github.com/pressly/goose/v3"
)

func (database *DB) RunMigrations() error {
	goose.SetBaseFS(migrations.FS)

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("set goose dialect: %w", err)
	}

	if err := goose.Up(database.db.DB, "."); err != nil {
		return fmt.Errorf("apply migrations: %w", err)
	}

	return nil
}
