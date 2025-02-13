package containers

import (
	"os"

	"github.com/abroudoux/dk/internal/logs"
	t "github.com/abroudoux/dk/internal/types"
	"github.com/charmbracelet/log"
)

func ContainerMode(ctx t.Context, cli t.Client, showAllContainers bool) error {
	containers, err := getContainers(ctx, cli, showAllContainers)
	if err != nil {
		return err
	}

	if len(containers) == 0 {
		logs.WarnMsg("No containers found.")
		return nil
	}

	containerSelected, err := selectContainer(containers)
	if err != nil {
		return err
	}

	if containerSelected.ID == "" {
		return nil
	}

	action, err := selectAction(containerSelected)
	if err != nil {
		return err
	}

	if action == containerActionExit {
		log.Info("Exiting program...")
		os.Exit(0)
	}

	err = doContainerAction(ctx, cli, containerSelected, action)
	if err != nil {
		return err
	}

	return nil
}
