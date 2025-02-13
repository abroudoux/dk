package containers

import (
	"fmt"

	"github.com/abroudoux/dk/internal/types"
	"github.com/charmbracelet/log"
)

func stopContainer(container Container, ctx types.Context, cli types.Client) error {
	stopOptions := StopOptions{}
	err := cli.ContainerStop(ctx, container.ID, stopOptions)
	if err != nil {
		return fmt.Errorf("error stopping container %s: %v", container.ID, err)
	}

	log.Info(fmt.Sprintf("Container %s stopped successfully", container.ID))
	return nil
}
