package containers

import (
	"fmt"

	t "github.com/abroudoux/dk/internal/types"
	"github.com/charmbracelet/log"
)

func stopContainer(container container, ctx t.Context, cli t.Client) error {
	stopOptions := stopOptions{}
	err := cli.ContainerStop(ctx, container.ID, stopOptions)
	if err != nil {
		return fmt.Errorf("error stopping container %s: %v", container.ID, err)
	}

	log.Info(fmt.Sprintf("Container %s stopped successfully", container.ID))
	return nil
}
