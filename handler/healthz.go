package handler

import (
	"encoding/json"
	"time"

	"log"

	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct {
}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	res := model.NewHealthzHandler("OK")
	time.Sleep(2 * time.Second)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
	}

}
