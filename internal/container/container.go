package container

import (
	"context"

	"github.com/docker/docker/api/types"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func GetContainers(ctx context.Context, cli *client.Client, allContainers bool) ([]types.Container, error) {
	containers, err := cli.ContainerList(ctx, containertypes.ListOptions{All: allContainers})
	if err != nil {
		return nil, err
	}
	return containers, err
}