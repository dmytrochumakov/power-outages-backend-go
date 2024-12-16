package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func write401Error(w http.ResponseWriter) {
	writeError(w, 401, "401 Unauthorized")
}

func writeError(w http.ResponseWriter, code int, error string) {
	type responseError struct {
		Error string `json:"error"`
	}
	respError := responseError{}

	w.WriteHeader(code)
	respError.Error = error

	dat, err := json.Marshal(respError)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		write500Error(w)
		return
	}
	w.Write(dat)
}

func write500Error(w http.ResponseWriter) {
	writeError(w, 500, "Something went wrong")
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, writeErr := w.Write(response)
	if writeErr != nil {
		log.Printf("Error writing response: %s", writeErr)
	}
}
