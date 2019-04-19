package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	var r *mux.Router = mux.NewRouter()
	r.HandleFunc("/", Handler)
	r.HandleFunc("/stores/{id}", Handler).Methods("GET")
	r.HandleFunc("/orders/", Handler).Methods("GET")
	r.Use(LoggingMiddleware)
	http.Handle("/", r)

	return r
}

func Handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jsonVars := []byte("OK")
	w.WriteHeader(http.StatusOK)

	jsonVars, _ = json.Marshal(vars)

	fmt.Fprintf(w, string(jsonVars))
}
