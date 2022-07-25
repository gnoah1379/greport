package vo

type ErrorResponse struct {
	Message string         `json:"message"`
	Detail  map[string]any `json:"detail"`
}
