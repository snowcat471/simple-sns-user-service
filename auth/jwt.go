package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

const minSecretSize = 20

type JWTMaker struct {
	secret string
}

func NewJWTMaker(secret string) (TokenMaker, error) {
	if len(secret) < minSecretSize {
		return nil, fmt.Errorf("secret length must be longer than %d", minSecretSize)
	}
	return &JWTMaker{secret}, nil
}

func (m *JWTMaker) CreateToken(payload *payload) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(m.secret))
}

func (m *JWTMaker) VerifyToken(token string) (*payload, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &payload{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(m.secret), nil
	})
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}
