package containers

import (
	"fmt"

	"github.com/abroudoux/dk/internal/types"
	"github.com/charmbracelet/log"
)

func pauseContainer(container Container, ctx types.Context, cli types.Client) error {
	err := cli.ContainerPause(ctx, container.ID)
	if err != nil {
		return fmt.Errorf("error pausing container %s: %v", container.ID, err)
	}

	log.Info(fmt.Sprintf("Container %s paused successfully", container.ID))
	return nil
}
