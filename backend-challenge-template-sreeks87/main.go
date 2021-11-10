package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sumup-challenges/backend-challenge-template-sreeks87/config"
	delivery "github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/delivery/http"
	"github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/repository"
	"github.com/sumup-challenges/backend-challenge-template-sreeks87/validator/service"
)

func main() {
	// The main function where all the init happens and the servers starts up

	// build the repository implementation
	dbPath := config.DB_FILE
	repo := repository.NewJSONFileRepository(dbPath)
	// build the service
	gen := service.NewFoodPositionService()
	svc := service.NewSnakeGameValidator(repo, gen)

	// build the handler
	r := mux.NewRouter()
	server := &http.Server{
		Handler:      r,
		Addr:         config.ADDRESS,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	delivery.NewValidatorController(r, svc)
	// start the sever

	log.Fatal(server.ListenAndServe())
}
