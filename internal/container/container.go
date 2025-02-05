package container

import (
	"context"

	"github.com/docker/docker/api/types"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func GetContainers(cli *client.Client, ctx context.Context) ([]types.Container, error) {
	containers, err := cli.ContainerList(ctx, containertypes.ListOptions{})
	if err != nil {
		return nil, err
	}
	return containers, err
}