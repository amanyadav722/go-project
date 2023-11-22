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

		next.ServeHTTP(w, r)
	})
}
