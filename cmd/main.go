package main

import (
	"context"
	"fmt"
	"os"

	"github.com/abroudoux/dk/internal/container"
	"github.com/abroudoux/dk/internal/forms"
	"github.com/abroudoux/dk/internal/logs"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		logs.Fatal("Error during cli init: ", err)
		os.Exit(1)
	}
	defer cli.Close()

	containers, err := container.GetContainers(cli, ctx)
	if err != nil {
		logs.Error("Error during containers recuperation: ", err)
		os.Exit(1)
	}

	containerSelected, err := forms.ChooseContainer(containers)
	if err != nil {
		logs.Error("Error during container selection: ", err)
	}

	action, err := forms.ChooseAction(containerSelected)
	if err != nil {
		logs.Error("Error during action selection: ", err)
	}

	fmt.Printf("%s", action)
}
