package api

type AuthUserService interface {
	Authorize(username string, password string) <-chan AuthResult
}

type AuthResult struct {
	Ok     bool
	UserId int64
}
