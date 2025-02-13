package containers

import (
	"github.com/abroudoux/dk/internal/types"
)

func doContainerAction(ctx types.Context, cli types.Client, container Container, action containerAction) error {
	switch action {
	case containerActionExit:
		return nil
	case containerActionCopyContainerID:
		return copyContainerId(container)
	case containerActionDelete:
		return deleteContainer(container, ctx, cli)
	case containerActionLogs:
		return getLogs(container, ctx, cli)
	case containerActionsStatus:
		status, err := selectStatus(container)
		if err != nil {
			return err
		}

		return doContainerStatusAction(container, status, ctx, cli)
	default:
		return nil
	}
}
