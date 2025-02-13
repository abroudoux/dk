package volumes

import (
	"context"

	"github.com/abroudoux/dk/internal/types"
	"github.com/charmbracelet/log"
	"github.com/docker/docker/client"
)

type Image = types.Image

func VolumeMode(ctx context.Context, cli *client.Client) error {
	log.Info("VolumeMode")
	return nil
}
