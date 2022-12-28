package database

import (
	"database/sql"
	"testing"
)

const connStr = "postgresql://user:password@localhost/place4live?sslmode=disable"
const migrations = "../../../migrations"

func TestOpenConnection(t *testing.T) {
	// GIVEN-WHEN
	db := testConnection(t)

	// THEN
	if db == nil {
		t.Errorf("Database connecton can't be nil!")
	}
}

func testConnection(t *testing.T) *sql.DB {
	conn, err := OpenConnection(connStr, migrations)
	if err != nil {
		t.Fatalf("Can't open database connection: connStr = %s, migrations = %s, error = %v\n", connStr, migrations, err)
	}
	return conn
}
