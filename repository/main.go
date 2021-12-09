package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sreeks87/repository/query/database"
	delivery "github.com/sreeks87/repository/query/delivery/http"
	"github.com/sreeks87/repository/query/downstream"
	"github.com/sreeks87/repository/query/service"
)

func main() {
	r := mux.NewRouter()
	server := http.Server{
		Handler: r,
	}
	m := make(map[string]int)
	db := database.NewKVMap(m)
	ds := downstream.NewBeChallenge("")
	svc := service.NewRepoSvc(db, ds)
	delivery.NewController(r, svc)
	log.Fatal(server.ListenAndServe())
}
