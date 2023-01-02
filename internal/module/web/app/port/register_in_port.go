package port

type RegisterInPort interface {
	Register(username string, password string) bool
}
