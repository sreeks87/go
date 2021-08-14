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

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
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

		id, _ := cmd.Flags().GetString("id")
		if id == "" {
			id = "all"
		}
		response, err := c.FetchTask(context.Background(), &tasks_grpc.TaskIDRequest{Id: id})

		DisplayOnConsole(response)
	},
}

func init() {
	taskCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	showCmd.PersistentFlags().StringP("id", "i", "", "ID of the task you want to view")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
