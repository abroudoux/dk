package docker

import (
	"context"

	"github.com/docker/docker/client"
)

func GetCtxCli() (context.Context, *client.Client, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return ctx, nil, err
	}
	defer cli.Close()

	return ctx, cli, nil
}