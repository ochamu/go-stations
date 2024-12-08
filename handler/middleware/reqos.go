package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mileusna/useragent"
)

type ContextKey string

const OSContextKey ContextKey = "os"

func DeviceOS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uaString := r.UserAgent()
		fmt.Println("UA: ", uaString)
		ua := useragent.Parse(uaString)
		os := ua.OS
		ctx := context.WithValue(r.Context(), OSContextKey, os)
		r = r.WithContext(ctx)

		h.ServeHTTP(w, r)
	})
}
