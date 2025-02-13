package containers

import (
	"fmt"

	t "github.com/abroudoux/dk/internal/types"
)

func doContainerStatusAction(container container, status containerStatus, ctx t.Context, cli t.Client) error {
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
