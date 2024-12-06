package middleware

import (
	"fmt"
	"net/http"
)

func Recovery(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// TODO: ここに実装をする
		defer func() {
			if rec := recover(); rec != nil {
				fmt.Println(rec)
				http.Error(w, "Server Error", http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// func LoggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		t1 := time.Now()
// 		next.ServeHTTP(w, r)
// 		t2 := time.Now()
// 		t := t2.Sub(t1)
// 		log.Printf("[%s] %s %s", r.Method, r.URL, t.String())
// 	})
// }

// func RecoverMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				debug.PrintStack()
// 				log.Printf("panic: %+v", err)
// 				http.Error(w, http.StatusText(500), 500)
// 			}
// 		}()
// 		next.ServeHTTP(w, r)
// 	})
// }
