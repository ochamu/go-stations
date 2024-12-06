package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	healthzH := handler.NewHealthzHandler()
	mux.HandleFunc("/healthz", healthzH.ServeHTTP)
	svc := service.NewTODOService(todoDB)
	todo := handler.NewTODOHandler(svc)
	mux.HandleFunc("/todos", todo.ServeHTTP)
	// mux.Handle("/do-panic", handler.NewDoPanicHandler())
	mux.Handle("/do-panic", middleware.Recovery(handler.NewDoPanicHandler()))
	return mux
}
