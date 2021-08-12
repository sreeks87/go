package main

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/schollz/jsonstore"
)

func main() {
	ks := new(jsonstore.JSONStore)

	// set a key to any object you want
	type Human struct {
		Name   string
		Height float64
	}
	err := ks.Set("human:0", Human{"Dante0", 5.4})
	if err != nil {
		panic(err)
	}

	// Saving will automatically gzip if .gz is provided,
	// and can be performed in a wait group
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = jsonstore.Save(ks, "test.json"); err != nil {
			panic(err)
		}
	}()
	wg.Wait()

	err = ks.Set("human:1", Human{"Dante1", 5.4})
	if err != nil {
		panic(err)
	}

	// Saving will automatically gzip if .gz is provided,
	// and can be performed in a wait group
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = jsonstore.Save(ks, "test.json"); err != nil {
			panic(err)
		}
	}()
	wg.Wait()

	err = ks.Set("human:2", Human{"Dante2", 5.4})
	if err != nil {
		panic(err)
	}

	// Saving will automatically gzip if .gz is provided,
	// and can be performed in a wait group
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = jsonstore.Save(ks, "test.json"); err != nil {
			panic(err)
		}
	}()
	wg.Wait()

	err = ks.Set("human:3", Human{"Dante3", 5.4})
	if err != nil {
		panic(err)
	}

	// Saving will automatically gzip if .gz is provided,
	// and can be performed in a wait group
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = jsonstore.Save(ks, "test.json"); err != nil {
			panic(err)
		}
	}()
	wg.Wait()

	// Load any JSON / GZipped JSON
	ks2, err := jsonstore.Open("test.json")
	if err != nil {
		panic(err)
	}

	// get the data back via an interface
	// var human Human
	var human Human
	err = ks2.Get("human:1", &human)
	all := ks2.GetAll(nil)
	for k, v := range all {
		log.Println(k)
		_ = json.Unmarshal(v, &human)
		log.Println(human.Name)
	}

	if err != nil {
		panic(err)
	}
}
