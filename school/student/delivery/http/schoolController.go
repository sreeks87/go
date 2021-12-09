package delivery

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sreeks87/school/student/domain"
)

type Controller struct {
	Service domain.Service
}

func NewController(r *mux.Router, svc domain.Service) {
	controller := &Controller{
		Service: svc,
	}
	r.HandleFunc("/student", controller.Register).Methods("POST")
	r.HandleFunc("/student/{id}", controller.GetStudent).Methods("GET")
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	req, _ := ioutil.ReadAll(r.Body)
	var std *domain.Student
	json.Unmarshal(req, &std)
	s, e := c.Service.Register(std)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(e)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

func (c *Controller) GetStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Println("params ", params)
	s, e := c.Service.Fetch(params["id"])
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(e)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}
