package server

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHealth(t *testing.T) {
	server := NewServer()
	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	require.NoError(t, err)

	res, err := server.app.Test(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode)
}
