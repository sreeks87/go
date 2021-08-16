package engine

import (
	"fmt"
	"net"
	grpcserver "pom/engine/task/controller/grpc"
	"pom/engine/task/infrastructure"
	"pom/engine/task/repository"

	"pom/engine/task/service"

	"google.golang.org/grpc"
	"gopkg.in/segmentio/analytics-go.v3"
)

const PORT = ":5001"

func Start(client analytics.Client, user string) {
	client.Enqueue(analytics.Track{
		UserId: user,
		Event:  "Starting server",
	})

	dbPath := "pomds.json"
	repo := repository.NewFileRepository(dbPath)
	segClient := infrastructure.NewLogger(client)
	svc := service.NewTaskService(repo, segClient, user)
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
		client.Enqueue(analytics.Track{
			UserId: user,
			Event:  "Error occured while starting server",
		})
	}

	client.Enqueue(analytics.Track{
		UserId: user,
		Event:  "Server Running",
	})
}
