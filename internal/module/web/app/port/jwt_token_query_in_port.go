package port

import "place4live/internal/module/web/domain"

type JwtTokenQueryInPort interface {
	Get(token string) (domain.JwtToken, error)
}
