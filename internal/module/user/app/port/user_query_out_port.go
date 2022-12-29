package port

import "place4live/internal/module/user/domain"

type UserQueryOutPort interface {
	GetUser(username string) <-chan *domain.User
}
