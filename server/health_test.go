package server

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestHealth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	server := NewServer(0)
	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	require.NoError(t, err)

	res, err := server.app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
}
