package engine

import (
	"fmt"
	"net"
	grpcserver "pom/engine/task/controller/grpc"
	"pom/engine/task/repository"

	"pom/engine/task/service"

	"google.golang.org/grpc"
)

const PORT = ":5001"

func Start() {
	dbPath := "pomds.json"
	repo := repository.NewFileRepository(dbPath)
	svc := service.NewTaskService(repo)
	list, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Could not listen.", err)
		panic(err)
	}

	gserver := grpc.NewServer()
	grpcserver.NewTaskServiceGRPC(gserver, svc)
	// fmt.Println("Server UP at: ", PORT)

	err = gserver.Serve(list)
	if err != nil {
		fmt.Println("Unexpected Error", err)
	}
}
