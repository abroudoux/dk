package compose

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func ComposeMode(ctx context.Context, cli *client.Client) error {
	err := composeBuild(ctx, cli)
	if err != nil {
		return fmt.Errorf("failed to build images: %w", err)
	}

	err = composeUp(ctx, cli)
	if err != nil {
		return fmt.Errorf("failed to start containers: %w", err)
	}

	return nil
}

func composeBuild(ctx context.Context, cli *client.Client) error {
	images := []string{"my_service_image"}

	for _, imageName := range images {
		fmt.Printf("Building image: %s\n", imageName)

		buildContext, err := os.Open("Dockerfile")
		if err != nil {
			return fmt.Errorf("could not open Dockerfile: %w", err)
		}
		defer buildContext.Close()

		_, err = cli.ImageBuild(ctx, buildContext, types.ImageBuildOptions{
			Tags:       []string{imageName},
			Remove:     true,
			ForceRemove: true,
		})
		if err != nil {
			return fmt.Errorf("failed to build image %s: %w", imageName, err)
		}
	}

	return nil
}

func composeUp(ctx context.Context, cli *client.Client) error {
	services := []string{"my_service"}

	for _, service := range services {
		fmt.Printf("Starting container for service: %s\n", service)

		resp, err := cli.ContainerCreate(ctx, &container.Config{
			Image: service,
		}, nil, nil, nil, service)

		if err != nil {
			return fmt.Errorf("could not create container for %s: %w", service, err)
		}

		err = cli.ContainerStart(ctx, resp.ID, containertypes.StartOptions{})
		if err != nil {
			return fmt.Errorf("could not start container %s: %w", service, err)
		}
	}

	return nil
}
