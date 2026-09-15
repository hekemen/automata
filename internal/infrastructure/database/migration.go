package database

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed migration.sql
var migrationSQL string

// RunMigrations executes the embedded migration.sql against the provided
// database pool. All statements are run in a single Exec call.
func RunMigrations(db *pgxpool.Pool) error {
	if _, err := db.Exec(context.Background(), migrationSQL); err != nil {
		return fmt.Errorf("execute migration: %w", err)
	}
	return nil
}
