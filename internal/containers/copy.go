package containers

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/log"
)

func copyContainerId(container Container) error {
	if clipboard.Unsupported {
		return fmt.Errorf("clipboard is unsupported on this platform.")
	}

	err := clipboard.WriteAll(container.ID)
	if err != nil {
		return fmt.Errorf("error copying container ID to clipboard: %v", err)
	}

	log.Info(fmt.Sprintf("Container ID copied to clipboard: %s", container.ID))
	return nil
}
