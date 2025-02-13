package containers

import (
	"fmt"
	"io"
	"os"

	t "github.com/abroudoux/dk/internal/types"
	"github.com/abroudoux/dk/internal/utils"
	"github.com/charmbracelet/log"
	containerTypes "github.com/docker/docker/api/types/container"
)

func getLogs(container container, ctx t.Context, cli t.Client) error {
	logOptions := containerTypes.LogsOptions{
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

	log.Info(fmt.Sprintf("Logs for container %s", utils.RenderContainerName(container)))

	_, err = io.Copy(os.Stdout, logsReader)
	if err != nil {
		return fmt.Errorf("error copying logs for container %s: %v", container.ID, err)
	}

	return nil
}
