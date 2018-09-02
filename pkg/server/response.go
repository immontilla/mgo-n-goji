package server

/** HTTP responses **/

import (
	"encoding/json"
	"net/http"
)

//Error http response
func Error(w http.ResponseWriter, code int, message string) {
	JSON(w, code, map[string]string{"error": message})
}

//JSON http response
func JSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//Code is an empty http response
func Code(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}
