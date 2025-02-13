package images

import (
	"fmt"

	t "github.com/abroudoux/dk/internal/types"
	"github.com/charmbracelet/log"
	"github.com/docker/docker/api/types/filters"
)

func prune(ctx t.Context, cli t.Client) error {
	_, err := cli.ImagesPrune(ctx, filters.NewArgs())
	if err != nil {
		return fmt.Errorf("failed to prune images: %v", err)
	}

	log.Info("Cleaned up dangling images")
	return nil
}
