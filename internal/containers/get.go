package containers

import (
	"fmt"

	t "github.com/abroudoux/dk/internal/types"
)

func getContainers(ctx t.Context, cli t.Client, showAllContainers bool) ([]container, error) {
	options := containerListOptions{
		All: showAllContainers,
	}
	containers, err := cli.ContainerList(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("error listing containers: %v", err)
	}
	return containers, nil
}
