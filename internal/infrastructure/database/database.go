package database

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

const driverPostgres = "postgres"
const statusCompleted = "completed"

func OpenConnection(connStr string, migrationsFile string) (*sql.DB, error) {
	db, err := sql.Open(driverPostgres, connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, runMigrations(db, migrationsFile)
}

func runMigrations(db *sql.DB, migrationsFile string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://"+migrationsFile, driverPostgres, driver)
	if err != nil {
		return err
	}
	err = m.Up()
	log.Printf("Database migration status: %s\n", migrationStatus(err))
	return nil
}

func migrationStatus(e error) string {
	if e == nil {
		return statusCompleted
	}
	return e.Error()
}
