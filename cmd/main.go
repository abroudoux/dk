package main

import (
	"context"
	"fmt"

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

	containerSelected, err := forms.ChooseContainer(containers)
	if err != nil {
		panic(err)
	}

	fmt.Println("%s", containerSelected.Names[0])
}
