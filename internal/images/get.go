package images

import (
	"fmt"

	t "github.com/abroudoux/dk/internal/types"
)

func getImages(ctx t.Context, cli t.Client, showAllImages bool) ([]image, error) {
	images, err := cli.ImageList(ctx, imageListOptions{All: showAllImages})
	if err != nil {
		return nil, fmt.Errorf("error during images recuperation: %v", err)
	}

	return images, err
}
