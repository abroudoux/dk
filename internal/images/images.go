package images

import (
	"context"
	"fmt"

	"github.com/abroudoux/dk/internal/logs"
	"github.com/abroudoux/dk/internal/types"
	"github.com/abroudoux/dk/internal/utils"
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

func DoImageAction(ctx context.Context, cli *client.Client, image Image, action ImageAction) error {
	switch action {
	case ImageActionExit:
		return nil
	case ImageActionDelete:
		return removeImage(image, ctx, cli)
	default:
		return nil
	}
}

func removeImage(image Image, ctx context.Context, cli *client.Client) error {
	removeOptions := imagetypes.RemoveOptions{}
	if _, err := cli.ImageRemove(ctx, image.ID, removeOptions); err != nil {
		return fmt.Errorf("error removing image %s: %v", image.ID, err)
	}

	logs.InfoMsg(fmt.Sprintf("Image %s removed", utils.RenderImageName(image)))
	return nil
}