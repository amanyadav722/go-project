package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	testHandler := AuthMiddleware(handler)

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", AuthToken)
	recorder := httptest.NewRecorder()
	testHandler.ServeHTTP(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)

}
