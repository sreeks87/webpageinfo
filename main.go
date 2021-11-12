package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sreeks87/webpageinfo/config"
)

// The mian function, eveything starts from here
func main() {
	// Creating a new router
	r := mux.NewRouter()

	// reading the template folder that hosts the ui files
	templateDir := http.Dir("./ui/")

	// handling url error with ui in url
	uiHandler := http.StripPrefix("/ui/", http.FileServer(templateDir))

	// handling the ui path prefix
	r.PathPrefix("/ui/").Handler(uiHandler).Methods("GET")

	// creating the server with minimum config
	// define the different timeouts below
	server := &http.Server{
		Handler:      r,
		Addr:         config.ADDRESS,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// handle the get for the UI
	r.HandleFunc("/ui", ui).Methods("GET")
	log.Fatal(server.ListenAndServe())
}

// the actual handler responsible for writing back to the UI
func ui(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "The ui is up...")
}
