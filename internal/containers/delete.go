package containers

import (
	"fmt"

	t "github.com/abroudoux/dk/internal/types"
	"github.com/abroudoux/dk/internal/utils"
	"github.com/charmbracelet/log"
)

func deleteContainer(container container, ctx t.Context, cli t.Client) error {
	removeOptions := containerRemoveOptions{
		Force:         true,
		RemoveVolumes: true,
	}
	if err := cli.ContainerRemove(ctx, container.ID, removeOptions); err != nil {
		return fmt.Errorf("error removing container %s: %v", container.ID, err)
	}

	log.Info(fmt.Sprintf("Container %s removed successfully", utils.RenderContainerName(container)))
	return nil
}
