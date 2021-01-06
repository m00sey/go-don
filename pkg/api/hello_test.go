package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGreetingHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	require.NoError(t, err)

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(greetingHandler)

	handler.ServeHTTP(rec, req)

	require.Equal(t, rec.Code, http.StatusOK)
	require.Equal(t, rec.Body.String(), "hi")
}
