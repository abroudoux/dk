package containers

import (
	"fmt"

	t "github.com/abroudoux/dk/internal/types"
	"github.com/charmbracelet/log"
)

func restartContainer(container container, ctx t.Context, cli t.Client) error {
	restartOptions := stopOptions{}
	err := cli.ContainerRestart(ctx, container.ID, restartOptions)
	if err != nil {
		return fmt.Errorf("error restarting container %s: %v", container.ID, err)
	}

	log.Info(fmt.Sprintf("Container %s restarted successfully", container.ID))
	return nil
}
