package container

import (
	"context"
	"fmt"

	"github.com/abroudoux/dk/internal/logs"
	"github.com/atotto/clipboard"
	"github.com/docker/docker/api/types"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Container = types.Container

func GetContainers(ctx context.Context, cli *client.Client, showAllContainers bool) ([]Container, error) {
	containers, err := cli.ContainerList(ctx, containertypes.ListOptions{All: showAllContainers})
	if err != nil {
		return nil, err
	}
	return containers, err
}

func DoContainerAction(containerSelected Container, action string) error {
	switch action {
	case "Copy Container ID":
		copyContainerId(containerSelected)
	}

	return nil
}

func copyContainerId(containerSelected Container) error {
	err := clipboard.WriteAll(containerSelected.ID)
	if err != nil {
		return err
	}
	logs.InfoMsg(fmt.Sprintf("Container ID copied to clipboard: %s", containerSelected.ID))
	return nil
}