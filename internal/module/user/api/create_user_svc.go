package api

type CreateUserService interface {
	Create(username string, password string) <-chan bool
}
