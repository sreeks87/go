/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"log"
	"pom/engine/task/controller/grpc/tasks_grpc"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(":5001", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		c := tasks_grpc.NewTaskHandlerClient(conn)

		desc, _ := cmd.Flags().GetString("desc")
		// response, err := c.(context.Background(), &chat.Message{Body: "Hello From Client!"})
		response, err := c.AddTask(context.Background(), &tasks_grpc.SingleTask{
			ID:          "0",
			Description: desc,
			State:       "Active",
		})
		ShowResponse(response, err)
	},
}

func init() {
	taskCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	addCmd.PersistentFlags().StringP("desc", "d", "", "Description of the task.")
	addCmd.MarkPersistentFlagRequired("desc")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("desc", "d", false, "Description of the task.")
}
