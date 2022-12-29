package test

import (
	"database/sql"
	"place4live/internal/lib/postgres"
	"testing"
)

const connStr = "postgresql://user:password@localhost/place4live?sslmode=disable"
const migrations = "../migrations"

func OpenDbConnection(t *testing.T) *sql.DB {
	conn, err := postgres.OpenConnection(connStr, migrations)
	if err != nil {
		t.Fatalf("Can't open database connection: connStr = %s, migrations = %s, error = %v\n", connStr, migrations, err)
	}
	t.Cleanup(func() {
		conn.Close()
	})
	return conn
}
