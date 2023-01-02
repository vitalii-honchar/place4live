package port

type CreateUserInPort interface {
	Create(username string, password string) <-chan bool
}
