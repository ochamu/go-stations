package middleware

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func BasicAuth(next http.Handler) http.Handler {
	err := godotenv.Load()
	if err != nil {
		// fmt.Errorf("Error")
	}
	userID := os.Getenv("BASIC_AUTH_USER_ID")
	password := os.Getenv("BASIC_AUTH_PASSWORD")
	// fmt.Println("PONPON", userID, password)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inputUserID, inputPassword, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if inputUserID != userID || inputPassword != password {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
