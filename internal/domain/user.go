package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id           int64
	Username     string
	PasswordHash string `json:"-"`
}

func NewUser(username string, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Username:     username,
		PasswordHash: string(hash),
	}, nil
}
