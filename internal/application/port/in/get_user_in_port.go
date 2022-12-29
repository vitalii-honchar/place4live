package in

import "place4live/internal/domain"

type GetUserInPort interface {
	GetUser(username string) <-chan *domain.User
}
