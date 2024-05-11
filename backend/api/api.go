package api

import (
	"encoding/json"
	"net/http"
)

type Payload struct {
    Meta json.RawMessage `json:"meta,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
}

type Response struct {
	Payload   Payload  `json:"payload"`
	Messages  []string `json:"messages"`
	ErrorCode int      `json:"errorCode"`
}
// NewResponse creates a new API response
func NewResponse(data interface{}, messages ...string) *Response {
	rawData, _ := json.Marshal(data) // Convert data to json.RawMessage
	return &Response{
		Payload:  Payload{Data: rawData},
		Messages: messages,
	}
}

// WriteJSON writes the JSON response to the http.ResponseWriter
func WriteJSON(w http.ResponseWriter, resp *Response, statusCode int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(resp)
}