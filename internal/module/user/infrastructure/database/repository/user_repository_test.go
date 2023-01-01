package repository

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"place4live/internal/module/user/domain"
	"place4live/test"
	"testing"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

const migrations = "../../../../../../migrations"

func TestPostgresUserRepository_Save(t *testing.T) {
	// GIVEN
	repository := NewUserRepository(test.OpenDbConnection(t, migrations))
	user := testUser(t)

	// WHEN
	actual := <-repository.Save(user)

	// THEN
	assert.True(t, actual)
}

func TestPostgresUserRepository_FindByUsername(t *testing.T) {
	t.Run("FindByUsername if user exists returns user", func(t *testing.T) {
		// GIVEN
		repository := NewUserRepository(test.OpenDbConnection(t, migrations))
		user := testUser(t)

		// WHEN
		<-repository.Save(user)
		actual := <-repository.FindByUsername(user.Username)

		// THEN
		assert.Equal(t, user.Username, actual.Username)
		assert.Equal(t, user.PasswordHash, actual.PasswordHash)
	})

	t.Run("FindByUsername if user doesn't exists returns nil", func(t *testing.T) {
		// GIVEN
		repository := NewUserRepository(test.OpenDbConnection(t, migrations))
		name := fmt.Sprintf("Username_%d", random.Intn(1000))

		// WHEN
		actual := <-repository.FindByUsername(name)

		// THEN
		assert.Nil(t, actual)
	})
}

func testUser(t *testing.T) *domain.User {
	u, err := domain.NewUser(
		fmt.Sprintf("Username_%d", random.Intn(1000)),
		fmt.Sprintf("Password_%d", random.Intn(1000)))
	if err != nil {
		t.Fatal(err)
	}
	return u
}
