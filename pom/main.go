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
package main

import (
	"math/rand"
	"os"
	"pom/cmd"
	"pom/engine"

	"gopkg.in/segmentio/analytics-go.v3"
)

func main() {
	WRITE_KEY := os.Getenv("WRITE_KEY")
	client := analytics.New(WRITE_KEY)
	defer client.Close()
	name := "testhost"
	invocationId := rand.Intn(100)
	user := "testuser"
	client.Enqueue(analytics.Identify{
		UserId: user,
		Traits: analytics.NewTraits().
			SetName(name).
			Set("invocationID", invocationId),
	})
	// Code to run the gRPC server
	// This code here should be as slim as possible
	// It should not do much logichere, abstract the logic to the server.

	// gRPC code start
	go engine.Start(client, user)
	cmd.Execute()
}
