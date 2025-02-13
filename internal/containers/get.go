package containers

import (
	"fmt"

	"github.com/abroudoux/dk/internal/types"
)

func getContainers(ctx types.Context, cli types.Client, showAllContainers bool) ([]Container, error) {
	options := types.ContainerListOptions{
		All: showAllContainers,
	}
	containers, err := cli.ContainerList(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("error listing containers: %v", err)
	}
	return containers, nil
}
