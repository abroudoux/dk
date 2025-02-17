package images

import (
	"github.com/abroudoux/dk/internal/logs"
	t "github.com/abroudoux/dk/internal/types"
	"github.com/charmbracelet/log"
)

func doImageAction(ctx t.Context, cli t.Client, image image, action imageAction) error {
	switch action {
	case imageActionExit:
		log.Info("Exiting..")
		return nil
	case imageActionDelete:
		return deleteImage(image, ctx, cli)
	case imageActionRun:
		return runImage(image, ctx, cli)
	default:
		logs.WarnMsg("Unknown action")
		return nil
	}
}
