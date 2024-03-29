package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		file, err := os.OpenFile("requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		logEntry := fmt.Sprintf("%s - %s %s\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)
		if _, err := file.WriteString(logEntry); err != nil {
			fmt.Println(err)
			return
		}

		recorder := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(recorder, r)

		if recorder.status >= 400 {
			logEntry := fmt.Sprintf("Error: %d - %s %s\n", recorder.status, r.Method, r.URL.Path)
			if _, err := file.WriteString(logEntry); err != nil {
				fmt.Println(err)
			}
		}
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.status = code
	r.ResponseWriter.WriteHeader(code)
}
