package auth

import (
	"math/rand"
	"testing"
	"time"

	"github.com/snowcat471/simple-sns-user-service/util"
	"github.com/stretchr/testify/assert"
)

func TestNewJWTMaker(t *testing.T) {
	testCases := []struct {
		name   string
		secret string
		test   func(t *testing.T, maker TokenMaker, err error)
	}{
		{
			name:   "PASS",
			secret: util.CreateRandomString(minSecretSize),
			test: func(t *testing.T, maker TokenMaker, err error) {
				assert.NotNil(t, maker)
				assert.NoError(t, err)
			},
		},
		{
			name:   "Invalid Secret Length",
			secret: util.CreateRandomString(minSecretSize - 1),
			test: func(t *testing.T, maker TokenMaker, err error) {
				assert.Nil(t, maker)
				assert.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			maker, err := NewJWTMaker(tc.secret)
			tc.test(t, maker, err)
		})
	}
}

func TestCreateToken(t *testing.T) {
	maker, _ := NewJWTMaker(util.CreateRandomString(rand.Intn(10) + minSecretSize))
	payload := NewPayload(rand.Intn(100)+1, time.Minute)

	token, err := maker.CreateToken(payload)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestVerifyToken(t *testing.T) {
	secret := util.CreateRandomString(rand.Intn(10) + minSecretSize)
	maker, _ := NewJWTMaker(secret)

	testCases := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "PASS",
			test: func(t *testing.T) {
				payload := NewPayload(rand.Intn(100)+1, time.Minute)
				token, _ := maker.CreateToken(payload)

				verifiedPayload, err := maker.VerifyToken(token)
				assert.NoError(t, err)
				assert.NotNil(t, verifiedPayload)
				assert.Equal(t, payload.UID, verifiedPayload.UID)
			},
		},
		{
			name: "Invalid Token",
			test: func(t *testing.T) {
				wrongToken := util.CreateRandomString(rand.Intn(100) + 1)

				invalidPayload, err := maker.VerifyToken(wrongToken)
				assert.EqualError(t, err, ErrInvalidToken.Error())
				assert.Nil(t, invalidPayload)
			},
		},
		{
			name: "Invalid Secret",
			test: func(t *testing.T) {
				wrongSecret := util.CreateRandomString(rand.Intn(10) + minSecretSize)
				wrongTokenMaker, _ := NewJWTMaker(wrongSecret)

				payload := NewPayload(rand.Intn(100)+1, time.Minute)
				token, _ := maker.CreateToken(payload)

				invalidPayload, err := wrongTokenMaker.VerifyToken(token)
				assert.EqualError(t, err, ErrInvalidToken.Error())
				assert.Nil(t, invalidPayload)
			},
		},
		{
			name: "Expired Token",
			test: func(t *testing.T) {
				payload := NewPayload(rand.Intn(100)+1, -time.Minute)
				token, _ := maker.CreateToken(payload)

				verifiedPayload, err := maker.VerifyToken(token)
				assert.EqualError(t, err, ErrExpiredToken.Error())
				assert.Nil(t, verifiedPayload)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.test(t)
		})
	}
}
