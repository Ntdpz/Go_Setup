package handlers

import (
	"net/http"
)

func HelloGolang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello golang"))
}
