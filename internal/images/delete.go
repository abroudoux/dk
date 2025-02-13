package images

import (
	"fmt"

	t "github.com/abroudoux/dk/internal/types"
	"github.com/abroudoux/dk/internal/utils"
	"github.com/charmbracelet/log"
)

func deleteImage(image image, ctx t.Context, cli t.Client) error {
	removeOptions := imageRemoveOptions{}
	if _, err := cli.ImageRemove(ctx, image.ID, removeOptions); err != nil {
		return fmt.Errorf("error removing image %s: %v", image.ID, err)
	}

	log.Info(fmt.Sprintf("Image %s removed", utils.RenderImageName(image)))
	return nil
}
