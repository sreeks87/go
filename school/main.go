package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	delivery "github.com/sreeks87/school/student/delivery/http"
	"github.com/sreeks87/school/student/domain"
	"github.com/sreeks87/school/student/repository"
	"github.com/sreeks87/school/student/service"
)

func main() {
	r := mux.NewRouter()
	server := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}
	db := make(map[string]domain.Student)
	repo := repository.NewKVMap(db)
	svc := service.NewStudentSvc(repo)
	delivery.NewController(r, svc)

	log.Fatal(server.ListenAndServe())
}
