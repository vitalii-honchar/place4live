package repository

import (
	"database/sql"
	"place4live/internal/lib"
	"place4live/internal/lib/postgres"
	"place4live/internal/module/user/domain"
)

const selectByUsername = "SELECT id, username, password_hash FROM p4l_user WHERE username = $1"
const insertUser = "INSERT INTO p4l_user (username, password_hash) VALUES ($1, $2)"

type UserRepository interface {
	FindByUsername(name string) <-chan *domain.User
	Save(user *domain.User) <-chan bool
}

type postgresUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &postgresUserRepository{db}
}

func (ur *postgresUserRepository) FindByUsername(name string) <-chan *domain.User {
	return lib.Async(func() *domain.User {
		return postgres.WithTx(ur.db, func(tx *sql.Tx) *domain.User {
			row := tx.QueryRow(selectByUsername, name)
			if row.Err() != nil {
				panic(row.Err())
			}
			var user domain.User
			if err := row.Scan(&user.Id, &user.Username, &user.PasswordHash); err != nil {
				panic(err)
			}
			return &user
		})
	})
}

func (ur *postgresUserRepository) Save(user *domain.User) <-chan bool {
	return lib.Async(func() bool {
		return postgres.WithTx(ur.db, func(tx *sql.Tx) bool {
			_, err := tx.Exec(insertUser, user.Username, user.PasswordHash)
			if err != nil {
				panic(err)
			}
			return true
		})
	})
}
