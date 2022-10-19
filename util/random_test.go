package util

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRandomString(t *testing.T) {
	length := rand.Intn(30) + 1
	randomString := CreateRandomString(length)
	assert.Equal(t, length, len(randomString))
}
