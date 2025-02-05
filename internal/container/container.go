package container

import (
	"context"
	"fmt"

	"github.com/abroudoux/dk/internal/forms"
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

func DoContainerAction(containerSelected Container, action forms.Action) error {
	switch action {
	case forms.ActionCopyContainerID:
		return copyContainerId(containerSelected)
	case forms.ActionExit:
		return nil
	default:
		return fmt.Errorf("unknown action: %v", action)
	}
}

func copyContainerId(containerSelected Container) error {
	err := clipboard.WriteAll(containerSelected.ID)
	if err != nil {
		return err
	}
	logs.InfoMsg(fmt.Sprintf("Container ID copied to clipboard: %s", containerSelected.ID))
	return nil
}