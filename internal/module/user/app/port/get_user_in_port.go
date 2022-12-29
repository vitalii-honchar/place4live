package port

import "place4live/internal/module/user/domain"

type GetUserInPort interface {
	GetUser(username string) <-chan *domain.User
}
