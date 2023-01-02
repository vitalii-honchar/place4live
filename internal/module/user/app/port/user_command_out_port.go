package port

import "place4live/internal/module/user/domain"

type UserCommandOutPort interface {
	Save(user *domain.User) <-chan bool
}
