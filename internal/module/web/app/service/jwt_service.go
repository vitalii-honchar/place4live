package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"place4live/internal/module/web/domain"
	"strconv"
	"time"
)

const claimAuthorized = "authorized"
const claimUserId = "user_id"
const claimExp = "exp"

type JwtService struct {
	secret string
}

func (js *JwtService) Get(token string) (domain.JwtToken, error) {
	jwtToken, err := js.convertToken(token)
	if err != nil || !jwtToken.Valid {
		return domain.JwtToken{}, err
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return domain.JwtToken{}, tokenError("unexpected JWT token")
	}
	userId, err := strconv.ParseInt(fmt.Sprintf("%d", claims[claimUserId]), 10, 64)
	if err != nil {
		return domain.JwtToken{}, err
	}
	expiredIn, err := strconv.ParseInt(fmt.Sprintf("%d", claims[claimUserId]), 10, 64)
	if err != nil {
		return domain.JwtToken{}, err
	}
	return domain.JwtToken{UserId: userId, ExpiredIn: time.UnixMilli(expiredIn)}, err
}

func (js *JwtService) generateToken(jt domain.JwtToken) (string, error) {
	claims := jwt.MapClaims{}
	claims[claimAuthorized] = true
	claims[claimUserId] = jt.UserId
	claims[claimExp] = jt.ExpiredIn.UnixMilli()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv(js.secret)))
}

func (js *JwtService) convertToken(t string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(js.secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

type tokenError string

func (t tokenError) Error() string {
	return string(t)
}
