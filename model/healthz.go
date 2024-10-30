package model

// A HealthzResponse expresses health check message.
type HealthzResponse struct {
	Message string `json:"message"`
}

func NewHealthzHandler(message string) *HealthzResponse {
	return &HealthzResponse{
		Message: message,
	}
}
