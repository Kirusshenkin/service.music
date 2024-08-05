package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", HealthCheckHandler).Methods("GET")
	return router
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
