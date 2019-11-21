package server

import "net/http"

// readinesHandler return proper success http status code
func readinesHandler(w http.ResponseWriter, r *http.Request) {
	writeStatusOK(w)
}

// livenessHandler return proper success http status code
func livenessHandler(w http.ResponseWriter, r *http.Request) {
	writeStatusOK(w)
}

func writeStatusOK(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}
