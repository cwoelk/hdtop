package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"os"
	"strings"
	"time"
)

func topLoop(client *docker.Client, topChan chan<- docker.TopResult) {
	for {
		topResult, _ := client.TopContainer(os.Args[1], strings.Join(os.Args[2:], " "))
		topChan <- topResult
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var (
		client *docker.Client
		err    error
	)

	if client, err = docker.NewClientFromEnv(); err != nil {
		fmt.Println("Could not connect to docker daemon:", err)
		os.Exit(1)
	}

	topChan := make(chan docker.TopResult)

	go topLoop(client, topChan)

	for {
		select {
		case topResult := <-topChan:
			for _, title := range topResult.Titles {
				fmt.Printf("%s\t", title)
			}
			fmt.Println()

			for _, process := range topResult.Processes {
				for _, value := range process {
					fmt.Printf("%s\t", value)
				}
				fmt.Println()
			}
		}
	}
}
