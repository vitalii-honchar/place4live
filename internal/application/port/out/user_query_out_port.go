package out

import "place4live/internal/domain"

type UserQueryOutPort interface {
	GetUser(username string) <-chan *domain.User
}
