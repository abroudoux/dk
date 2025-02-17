package images

import (
	t "github.com/abroudoux/dk/internal/types"
	"github.com/charmbracelet/log"
)

func ImageMode(ctx t.Context, cli t.Client) error {
	images, err := getImages(ctx, cli, false)
	if err != nil {
		return err
	}

	imageSelected, err := selectImage(images)
	if err != nil {
		return err
	}

	if imageSelected.ID == "" {
		log.Warn("Program exited without selecting an image..")
		return nil
	}

	action, err := selectAction(imageSelected)
	if err != nil {
		return err
	}

	err = doImageAction(ctx, cli, imageSelected, action)
	if err != nil {
		return err
	}

	return nil
}

func BuildMode(ctx t.Context, cli t.Client) error {
	err := buildImage(ctx, cli)
	if err != nil {
		return err
	}

	return nil
}
