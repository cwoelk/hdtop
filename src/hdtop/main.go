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
	fmt.Printf("%#v", topResult)
}
