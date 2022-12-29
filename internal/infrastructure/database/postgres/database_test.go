package postgres

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const connStr = "postgresql://user:password@localhost/place4live?sslmode=disable"
const migrations = "../../../../migrations"

func TestOpenConnection(t *testing.T) {
	// GIVEN-WHEN
	conn, err := OpenConnection(connStr, migrations)
	t.Cleanup(func() {
		if conn != nil {
			conn.Close()
		}
	})

	// THEN
	assert.Nil(t, err)
}
