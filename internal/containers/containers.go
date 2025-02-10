package containers

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/abroudoux/dk/internal/logs"
	"github.com/abroudoux/dk/internal/types"
	"github.com/abroudoux/dk/internal/utils"
	"github.com/atotto/clipboard"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Container = types.Container

func ContainerMode(ctx context.Context, cli *client.Client, showAllContainers bool) error {
	containers, err := getContainers(ctx, cli, showAllContainers)
	if err != nil {
		return err
	}

	if len(containers) == 0 {
		logs.WarnMsg("No containers found.")
		return nil
	}

	sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() error {
		containerSelected, err := selectContainer(containers)
		if err != nil {
			return err
		}

		if containerSelected.ID == "" {
			return nil
		}

		action, err := selectAction(containerSelected)
		if err != nil {
			return err
		}

		err = doContainerAction(ctx, cli, containerSelected, action)
		if err != nil {
			return err
		}

		return nil
    }()

	<-sigChan
    fmt.Println("\nProgram interrupted. Exiting...")
    os.Exit(0)

	return nil
}

func getContainers(ctx context.Context, cli *client.Client, showAllContainers bool) ([]Container, error) {
	options := types.ContainerListOptions{
		All: showAllContainers,
	}
	containers, err := cli.ContainerList(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("error listing containers: %v", err)
	}
	return containers, nil
}

func doContainerAction(ctx context.Context, cli *client.Client, container Container, action ContainerAction) error {
	switch action {
	case ContainerActionExit:
		return nil
	case ContainerActionCopyContainerID:
		return copyContainerId(container)
	case ContainerActionDelete:
		return deleteContainer(container, ctx, cli)
	case ContainerActionLogs:
		return getLogs(container, ctx, cli)
	case ContainerActionsStatus:
		return getStatus(container, ctx, cli)
	default:
		return nil
	}
}

func copyContainerId(container Container) error {
	err := clipboard.WriteAll(container.ID)
	if err != nil {
		return fmt.Errorf("error copying container ID to clipboard: %v", err)
	}
	logs.InfoMsg(fmt.Sprintf("Container ID copied to clipboard: %s", container.ID))
	return nil
}

func getStatus(container Container, ctx context.Context, cli *client.Client) error {
	status, err := SelectStatus(container)
	if err != nil {
		return err
	}

	switch status {
	case ContainerStatusExit:
		return nil
	case ContainerStatusPause:
		return pauseContainer(container, ctx, cli)
	case ContainerStatusRestart:
		return restartContainer(container, ctx, cli)
	case ContainerStatusStop:
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

	logs.InfoMsg(fmt.Sprintf("Container %s removed successfully", utils.RenderContainerName(container)))
    return nil
}

func getLogs(container Container, ctx context.Context, cli *client.Client) error {
    logOptions := containertypes.LogsOptions{
        ShowStdout: true,
        ShowStderr: true,
        Follow:     true,
        Timestamps: true,
        Tail:       "50",
    }

    logsReader, err := cli.ContainerLogs(ctx, container.ID, logOptions)
    if err != nil {
        return fmt.Errorf("error getting logs for container %s: %v", container.ID, err)
    }
    defer logsReader.Close()

    logs.InfoMsg(fmt.Sprintf("Logs for container %s", utils.RenderContainerName(container)))

    _, err = io.Copy(os.Stdout, logsReader)
    if err != nil {
        return fmt.Errorf("error copying logs for container %s: %v", container.ID, err)
    }

    return nil
}

