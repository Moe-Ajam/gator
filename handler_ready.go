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

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
