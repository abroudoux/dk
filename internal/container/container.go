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

func DoContainerAction(ctx context.Context, cli *client.Client, container Container, action forms.Action) error {
	switch action {
	case forms.ActionExit:
		return nil
	case forms.ActionCopyContainerID:
		return copyContainerId(container)
	case forms.ActionDelete:
		return deleteContainer(container, ctx, cli)
	case forms.ActionsStatus:
		return getStatus(container, ctx, cli)
	default:
		return fmt.Errorf("unknown action: %v", action)
	}
}

func copyContainerId(container Container) error {
	err := clipboard.WriteAll(container.ID)
	if err != nil {
		return err
	}
	logs.InfoMsg(fmt.Sprintf("Container ID copied to clipboard: %s", container.ID))
	return nil
}

func getStatus(container Container, ctx context.Context, cli *client.Client) error {
	status, err := forms.ChooseStatus(container)
	if err != nil {
		logs.ErrorMsg(fmt.Sprintf("Error choosing status: %v", err))
		return err
	}

	switch status {
	case forms.StatusExit:
		return nil
	case forms.StatusPause:
		return pauseContainer(container, ctx, cli)
	case forms.StatusRestart:
		return restartContainer(container, ctx, cli)
	case forms.StatusStop:
		return stopContainer(container, ctx, cli)
	default:
		return fmt.Errorf("unknown status: %v", status)
	}
}

func pauseContainer(container Container, ctx context.Context, cli *client.Client) error {
	err := cli.ContainerPause(ctx, container.ID)
	if err != nil {
		return fmt.Errorf("error pausing container %s: %v", container.ID, err)
	}
	logs.InfoMsg(fmt.Sprintf("Container %s paused successfully", container.ID))
	return nil
}

func restartContainer(container Container, ctx context.Context, cli *client.Client) error {
	restartOptions := containertypes.StopOptions{}
	err := cli.ContainerRestart(ctx, container.ID, restartOptions)
	if err != nil {
		return fmt.Errorf("error restarting container %s: %v", container.ID, err)
	}
	logs.InfoMsg(fmt.Sprintf("Container %s restarted successfully", container.ID))
	return nil
}

func stopContainer(container Container, ctx context.Context, cli *client.Client) error {
	stopOptions := containertypes.StopOptions{}
	err := cli.ContainerStop(ctx, container.ID, stopOptions)
	if err != nil {
		return fmt.Errorf("error stopping container %s: %v", container.ID, err)
	}
	logs.InfoMsg(fmt.Sprintf("Container %s stopped successfully", container.ID))
	return nil
}

func deleteContainer(container Container, ctx context.Context, cli *client.Client) error {
	removeOptions := containertypes.RemoveOptions{
        Force: true,
        RemoveVolumes: true,
    }
	if err := cli.ContainerRemove(ctx, container.ID, removeOptions); err != nil {
        return fmt.Errorf("error removing container %s: %v", container.ID, err)
    }

	logs.InfoMsg(fmt.Sprintf("Container %s removed successfully", container.ID))
    return nil
}