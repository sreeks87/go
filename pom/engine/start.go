package engine

import (
	"controller/grpc/"
	"log"
	"net"

	"github.com/golang/protobuf/protoc-gen-go/grpc"
)

func start() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := tasks_grpc.Server()
	grpcServer := grpc.NewServer()
	tasks_grpc.RegisterTaskHandlerServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start: %s", err)
	}

}
