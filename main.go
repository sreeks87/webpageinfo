package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sreeks87/webpageinfo/config"
	delivery "github.com/sreeks87/webpageinfo/pageinfo/delivery/http"
)

// The main function, eveything starts from here
func main() {
	// Creating a new router via newRouter function

	r := newRouter()
	// creating the server with minimum config
	// define the minimum server config
	server := &http.Server{
		Handler:      r,
		Addr:         config.ADDRESS,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

}

func newRouter() *mux.Router {
	// using mux router here and creating a new router
	r := mux.NewRouter()

	// handling the routes for the api
	// there are two endpoints /webpageinfo that shows the details and /heartbeat that acts
	// as a healthcheck api
	// the /webpageinfo api is consumed in the html page for getting page info
	r.StrictSlash(true)
	r.HandleFunc("/webpageinfo", delivery.Webpageinfo).Methods("POST")
	r.HandleFunc("/heartbeat", delivery.Heartbeat).Methods("GET")

	// reading the template folder that hosts the ui files
	// the html and js files are in thisfolder, using it as the static dir.
	templateDir := http.Dir("./ui/")

	// serving the static files
	// when a route sets a path prefix using the PathPrefix() method,
	// strict slash is ignored
	uiHandler := http.StripPrefix("/ui/", http.FileServer(templateDir))
	r.PathPrefix("/ui/").Handler(uiHandler).Methods("GET")

	return r
}
