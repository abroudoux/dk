package containers

import (
	"fmt"

	t "github.com/abroudoux/dk/internal/types"
	"github.com/charmbracelet/log"
)

func pauseContainer(container container, ctx t.Context, cli t.Client) error {
	err := cli.ContainerPause(ctx, container.ID)
	if err != nil {
		return fmt.Errorf("error pausing container %s: %v", container.ID, err)
	}

	log.Info(fmt.Sprintf("Container %s paused successfully", container.ID))
	return nil
}
