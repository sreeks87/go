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
	"fmt"
	"log"
	"pom/engine/task/controller/grpc/tasks_grpc"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		taskLimit, breakLimit, bigBreakLimit := getRunConfig()
		RunTaskinLoop(id, taskLimit, breakLimit, bigBreakLimit)
	},
}

func init() {
	taskCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	runCmd.PersistentFlags().StringP("id", "i", "", "ID of the task you want to run")
	runCmd.MarkPersistentFlagRequired("id")
}

func RunTaskinLoop(id string, taskLimit int64, breakLimit int64, bigBreakLimit int64) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := tasks_grpc.NewTaskHandlerClient(conn)
	response, _ := c.FetchTask(context.Background(), &tasks_grpc.TaskIDRequest{Id: id})
	desc := response.Tasks[0].Description
	// todo a pause feature to pause the running task.
	// todo use ui-contols libraries to make it look better.
	// https://github.com/avelino/awesome-go#advanced-console-uis
	count := 0
	for {
		fmt.Println("Press Ctrl C to quit.")
		var barTask Bar
		count += 1
		barTask.NewOption(0, taskLimit)

		fmt.Println("Executing task : '", desc, "'. Instance :", count)
		Loop(taskLimit, barTask)
		barTask.Finish()

		beep()
		var barBreak Bar
		breakMsg := "Break time :) "
		if count == 4 {
			breakLimit = bigBreakLimit
			breakMsg = "Take a long break.You deserve it!!!"
		}
		barBreak.NewOption(0, breakLimit)
		fmt.Println(breakMsg)

		Loop(breakLimit, barBreak)
		barBreak.Finish()
		beep()
	}
}

func Loop(stop int64, bar Bar) {
	for i := 0; int64(i) <= stop; i++ {
		time.Sleep(1 * time.Second)
		bar.Play(int64(i))
	}
}

func getRunConfig() (int64, int64, int64) {
	// todo read these from the config
	return 25 * 60, 5 * 60, 20 * 60
}

func beep() {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		panic(err)
	}
}
