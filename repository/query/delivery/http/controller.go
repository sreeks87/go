package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sreeks87/repository/query/domain"
)

// controller --> service ---> donstream/database

type Controller struct {
	Svc domain.Service
}

func NewController(r *mux.Router, s domain.Service) {
	controller := &Controller{
		Svc: s,
	}

	r.HandleFunc("/repositories", controller.GetRepositories).Methods("GET")
	r.HandleFunc("/statistics", controller.GetStat).Methods("GET")
}

func (c *Controller) GetRepositories(w http.ResponseWriter, r *http.Request) {
	queryparam := r.URL.Query().Get("query")
	res, e := c.Svc.Fetch(queryparam)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(e)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}

func (c *Controller) GetStat(w http.ResponseWriter, r *http.Request) {
	res, e := c.Svc.Stat()
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(e)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}
