package models

// HTTPResponse struct
type HTTPResponse struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}
