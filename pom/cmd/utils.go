package cmd

import (
	"fmt"
	"log"
	"pom/engine/task/controller/grpc/tasks_grpc"

	"google.golang.org/grpc"
)

func GetGRPCTaskHandler() tasks_grpc.TaskHandlerClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := tasks_grpc.NewTaskHandlerClient(conn)

	return c
}

func ShowResponse(response interface{}, err error) {
	if err != nil {
		log.Fatalf("Error when calling AddTask: %s", err)
	}
	fmt.Println("<-------------------------------------------------------------------------------------------->")
	fmt.Println(response)
	fmt.Println("<-------------------------------------------------------------------------------------------->")
}

func DisplayOnConsole(response *tasks_grpc.ListTasks) {
	fmt.Println("<-------------------------------------------------------------------------------------------->")
	for _, v := range response.Tasks {
		fmt.Println(v.ID, "\t", v.Description, "\t", v.State)
	}
	fmt.Println("<-------------------------------------------------------------------------------------------->")
}

// func ShowMultiResponse(err error, response *tasks_grpc.ListTasks) {
// 	if err != nil {
// 		log.Fatalf("Error when calling AddTask: %s", err)
// 	}
// 	log.Printf("Response from server: %s", response)
// }
