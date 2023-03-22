package api

import (
	"encoding/json"
	"fmt"
)

type APIErrorMessage struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	Suggestion string `json:"suggestion,omitempty"`
}

type APIError struct {
	APIErr      APIErrorMessage            `json:"error"`
	FieldErrors map[string]APIErrorMessage `json:"field_errors"`
}

func (e *APIError) Raw() ([]byte, error) {
	return json.Marshal(e)
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error got code: %s, message: %s, suggestion: %s", e.APIErr.Code, e.APIErr.Message, e.APIErr.Suggestion)
}
