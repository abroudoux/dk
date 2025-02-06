package images

import (
	"context"

	"github.com/abroudoux/dk/internal/types"
	imagetypes "github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

type Image = types.Image

func GetImages(ctx context.Context, cli *client.Client, showAllImages bool) ([]Image, error) {
	images, err := cli.ImageList(ctx, imagetypes.ListOptions{All: showAllImages})
	if err != nil {
		return nil, err
	}
	return images, err
}