package utils

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func NewULID() (ulid.ULID, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	return ulid.New(ms, entropy)
}
