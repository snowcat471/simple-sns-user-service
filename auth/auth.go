package auth

import (
	"errors"
	"time"
)

var (
	ErrExpiredToken = errors.New("expired token")
	ErrInvalidToken = errors.New("invalid token")
)

type TokenMaker interface {
	CreateToken(payload *payload) (string, error)
	VerifyToken(token string) (*payload, error)
}

type payload struct {
	UID       int       `json:"uid"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(UID int, duration time.Duration) *payload {
	return &payload{
		UID:       UID,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}

func (p *payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
