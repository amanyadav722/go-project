package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoggerMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	testHandler := Logger(handler)

	req, _ := http.NewRequest("GET", "/testpath", nil)
	recorder := httptest.NewRecorder()
	testHandler.ServeHTTP(recorder, req)

}
