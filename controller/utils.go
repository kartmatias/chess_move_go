package controller

import "net/http"

func ReturnError(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.WriteHeader(errorCode)
	w.Write([]byte(errorMsg))
}
