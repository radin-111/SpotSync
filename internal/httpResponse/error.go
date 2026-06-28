package httpResponse

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Errors  string `json:"errors,omitempty"`
}
