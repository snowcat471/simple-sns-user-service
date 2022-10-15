package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewULID(t *testing.T) {
	u1, err := NewULID()
	assert.NoError(t, err)
	u2, err := NewULID()
	assert.NoError(t, err)
	assert.NotEqual(t, u1, u2)
}
