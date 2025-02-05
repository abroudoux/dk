package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/abroudoux/dk/internal/container"
	"github.com/abroudoux/dk/internal/forms"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containers, err := container.GetContainers(cli, ctx)
	if err != nil {
		panic(err)
	}

	for _, c := range containers {
		fmt.Printf("Name: %s \n", strings.Join(c.Names, ""))
	}

	containerSelected, err := forms.ChooseContainer(containers)
	if err != nil {
		panic(err)
	}

	// containerNameSelected := strings.Join(containerSelected.Names, "")
	// fmt.Printf("%s", containerNameSelected)

	action, err := forms.ChooseAction(containerSelected)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", action)
}
