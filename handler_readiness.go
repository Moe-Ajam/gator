package main

import "net/http"

type readinessResponse struct {
	Status string `json:"status"`
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, readinessResponse{
		Status: "ok",
	})
}
