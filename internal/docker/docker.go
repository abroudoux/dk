package docker

import (
	"fmt"

	"github.com/docker/docker/client"
)

func GetCli() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("error during docker client initialization: %w", err)
	}
	defer cli.Close()

	return cli, nil
}