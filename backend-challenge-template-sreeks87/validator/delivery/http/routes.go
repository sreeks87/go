package delivery

import "github.com/gorilla/mux"

func Route(r *mux.Router, controller *ValidatorController) {
	r.HandleFunc("/new", controller.GetNewGame).Methods("GET")
	r.HandleFunc("/validate", controller.Validate).Methods("POST")
}
