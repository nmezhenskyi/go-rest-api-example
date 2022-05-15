package webserver

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error      string   `json:"error"`
	Message    string   `json:"message,omitempty"`
	Details    []string `json:"details,omitempty"`
	StatusCode int      `json:"statusCode"`
}

func sendJSON(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if body != nil {
		json.NewEncoder(w).Encode(body)
	} else {
		json.NewEncoder(w).Encode(struct{}{})
	}
}

func sendBadRequest(w http.ResponseWriter, message string, details []string) {
	err := errorResponse{
		Error:      "Bad Request",
		Message:    message,
		Details:    details,
		StatusCode: http.StatusBadRequest,
	}
	sendJSON(w, http.StatusBadRequest, err)
}

func sendNotFound(w http.ResponseWriter, message string) {
	err := errorResponse{
		Error:      "Not Found",
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
	sendJSON(w, http.StatusNotFound, err)
}

func sendInternalError(w http.ResponseWriter) {
	err := errorResponse{
		Error:      "Internal Server Error",
		StatusCode: http.StatusInternalServerError,
	}
	sendJSON(w, http.StatusInternalServerError, err)
}
