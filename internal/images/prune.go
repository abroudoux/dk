package images

import (
	"context"
	"fmt"

	"github.com/abroudoux/dk/internal/logs"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func PruneImages(ctx context.Context, cli *client.Client) error {
    _, err := cli.ImagesPrune(ctx, filters.NewArgs())
    if err != nil {
        return fmt.Errorf("failed to prune images: %v", err)
    }
    logs.InfoMsg("Cleaned up dangling images")
    return nil
}