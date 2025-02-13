package containers

import (
	"fmt"

	"github.com/abroudoux/dk/internal/types"
)

func doContainerStatusAction(container Container, status containerStatus, ctx types.Context, cli types.Client) error {
	switch status {
	case containerStatusExit:
		return nil
	case containerStatusPause:
		return pauseContainer(container, ctx, cli)
	case containerStatusRestart:
		return restartContainer(container, ctx, cli)
	case containerStatusStop:
		return stopContainer(container, ctx, cli)
	default:
		return fmt.Errorf("unknown status: %v", status)
	}
}
