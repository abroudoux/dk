package volumes

import (
	"context"

	"github.com/abroudoux/dk/internal/types"
	"github.com/docker/docker/client"
)

type Image = types.Image

func VolumeMode(ctx context.Context, cli *client.Client) error {
	return nil
}
