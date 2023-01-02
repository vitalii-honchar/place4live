package domain

import "time"

const tokenLifeSpan = 1 * time.Hour

type JwtToken struct {
	UserId    int64
	ExpiredIn time.Time
}

func NewJwtToken(userId int64) JwtToken {
	return JwtToken{UserId: userId, ExpiredIn: time.Now().Add(tokenLifeSpan)}
}

func (t JwtToken) IsValid() bool {
	return time.Now().Add(-1 * tokenLifeSpan).Before(t.ExpiredIn)
}
