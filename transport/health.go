package transport

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method", 405)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
