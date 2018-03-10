package service

import "net/http"

func ResponseBadRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("400 - Bad Request"))
}

func ResponseNotFound(w http.ResponseWriter, t string) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - " + t))
}
