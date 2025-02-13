package containers

import (
	"fmt"

	"github.com/abroudoux/dk/internal/types"
	"github.com/abroudoux/dk/internal/utils"
	"github.com/charmbracelet/log"
	containerTypes "github.com/docker/docker/api/types/container"
)

func deleteContainer(container Container, ctx types.Context, cli types.Client) error {
	removeOptions := containerTypes.RemoveOptions{
		Force:         true,
		RemoveVolumes: true,
	}
	if err := cli.ContainerRemove(ctx, container.ID, removeOptions); err != nil {
		return fmt.Errorf("error removing container %s: %v", container.ID, err)
	}

	log.Info(fmt.Sprintf("Container %s removed successfully", utils.RenderContainerName(container)))
	return nil
}
