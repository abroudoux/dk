package containers

import (
	"fmt"

	"github.com/abroudoux/dk/internal/types"
	"github.com/charmbracelet/log"
)

func restartContainer(container Container, ctx types.Context, cli types.Client) error {
	restartOptions := StopOptions{}
	err := cli.ContainerRestart(ctx, container.ID, restartOptions)
	if err != nil {
		return fmt.Errorf("error restarting container %s: %v", container.ID, err)
	}

	log.Info(fmt.Sprintf("Container %s restarted successfully", container.ID))
	return nil
}
