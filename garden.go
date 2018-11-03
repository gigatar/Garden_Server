package main

import (
	rest "Garden/rest/sensors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	restEndpoint = "/rest/v1.0/"
)

func main() {
	// Setup our router for handling requests
	// Could use strictslaches and milddlewhere here too.
	router := mux.NewRouter()

	restAPI := router.PathPrefix(restEndpoint).Subrouter()
	// Health functions
	restAPI.HandleFunc("/sensor", rest.AddSensorData).Methods("POST")

	// User functions
	// rest.HandleFunc("/users", users.ListUsers).Methods("GET")

	// UI
	staticFileDirectory := http.Dir("./web/")
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
	router.PathPrefix("/").Handler(staticFileHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
