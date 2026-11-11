package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	_ = respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
