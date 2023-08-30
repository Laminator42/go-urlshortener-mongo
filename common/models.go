package common

// TODO: Not the cleanest solution? Can I import that without common. prefix?
type ErrorResponse struct {
	Error string `json:"error"`
}
