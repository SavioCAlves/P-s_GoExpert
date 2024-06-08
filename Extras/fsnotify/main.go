package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

type DBconfig struct {
	DB       string `json:"db"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var config DBconfig

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()
	MarshallConfig("config.json")

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println("event: ", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					MarshallConfig("config.json")
					fmt.Println("modified file: ", event.Name)
					fmt.Println(config)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error: ", err)
			}
		}
	}()

	err = watcher.Add("config.json")
	if err != nil {
		panic(err)
	}
	<-done
}

func MarshallConfig(file string) {
	date, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(date, &config)
	if err != nil {
		panic(err)
	}
}
