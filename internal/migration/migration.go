package migration

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

// Runner handles database migrations
type Runner struct {
	db            *gorm.DB
	migrationsDir string
}

// NewRunner creates a new migration runner
func NewRunner(db *gorm.DB, migrationsDir string) *Runner {
	return &Runner{
		db:            db,
		migrationsDir: migrationsDir,
	}
}

// Run executes all pending migrations
func (r *Runner) Run() error {
	// Get the underlying SQL database
	sqlDB, err := r.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB from gorm: %w", err)
	}

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", r.migrationsDir),
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	// Run migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	if err == migrate.ErrNoChange {
		log.Println("No new migrations to apply")
	} else {
		log.Println("Migrations applied successfully")
	}

	return nil
}

// Version returns the current migration version
func (r *Runner) Version() (uint, bool, error) {
	// Get the underlying SQL database
	sqlDB, err := r.db.DB()
	if err != nil {
		return 0, false, fmt.Errorf("failed to get sql.DB from gorm: %w", err)
	}

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return 0, false, fmt.Errorf("failed to create migration driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", r.migrationsDir),
		"postgres",
		driver,
	)
	if err != nil {
		return 0, false, fmt.Errorf("failed to create migrate instance: %w", err)
	}

	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return 0, false, fmt.Errorf("failed to get migration version: %w", err)
	}

	return version, dirty, nil
}

// CheckAndRun checks if migrations are needed and runs them if necessary
func (r *Runner) CheckAndRun() error {
	version, dirty, err := r.Version()
	if err != nil && err.Error() != "failed to get migration version: no migration" {
		return fmt.Errorf("failed to check migration status: %w", err)
	}

	if dirty {
		return fmt.Errorf("database is in dirty state at version %d, manual intervention required", version)
	}

	log.Printf("Current migration version: %d", version)

	// Run migrations
	return r.Run()
}
