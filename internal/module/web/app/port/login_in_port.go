package port

type LoginInPort interface {
	Login(username string, password string) (string, bool)
}
