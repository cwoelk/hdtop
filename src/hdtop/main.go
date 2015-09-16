package main

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"os"
	"strings"
)

func main() {
	var (
		client *docker.Client
		err    error
	)

	if client, err = docker.NewClientFromEnv(); err != nil {
		fmt.Println("Could not connect to docker daemon:", err)
		os.Exit(1)
	}

	topResult, _ := client.TopContainer(os.Args[1], strings.Join(os.Args[2:], " "))

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
