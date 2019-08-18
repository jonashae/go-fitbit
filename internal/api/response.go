package api

// Error represents an error message returned from the API
type Error struct {
	Type    string `json:"errorType"`
	Message string `json:"message"`
}

// GenericResponse represents response returned from the API
// that is not associated with a specific entity type
type GenericResponse struct {
	Success bool    `json:"success"`
	Errors  []Error `json:"errors"`
}
