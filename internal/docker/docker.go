package docker

import (
	"fmt"

	"github.com/abroudoux/dk/internal/types"
	"github.com/docker/docker/client"
)

func GetClient() (types.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("error during docker client initialization: %w", err)
	}
	defer cli.Close()

	return cli, nil
}
