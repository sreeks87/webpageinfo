package delivery

import "github.com/gorilla/mux"

func route(r *mux.Router) {
	r.HandleFunc("/webpageinfo", webpageinfo).Methods("POST")
	r.HandleFunc("/heartbeat", heartbeat).Methods("GET")
}
