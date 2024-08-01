package models

// Response model info
// @Description Response information
type Response struct {
	Message string `json:"message,omitempty"` // Response message
	Error   string `json:"error,omitempty"`
}
