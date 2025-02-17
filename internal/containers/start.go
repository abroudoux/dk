package containers

import (
	t "github.com/abroudoux/dk/internal/types"
)

func StartContainer(ctx t.Context, cli t.Client, containerId string, options ContainerStartOptions) error {
	err := cli.ContainerStart(ctx, containerId, options)
	if err != nil {
		return err
	}

	return nil
}
