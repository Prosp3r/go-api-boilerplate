package response

// HTTPError allows you yo return nice error responses
type HTTPError struct {
	Code    int
	Error   error
	Message string `json:"message"`
}
