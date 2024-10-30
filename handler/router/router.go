package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	healthzH := handler.NewHealthzHandler()
	mux.HandleFunc("/healthz", healthzH.ServeHTTP)
	return mux
}
