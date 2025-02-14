package history

import (
	"os"

	"github.com/charmbracelet/log"
)

func readHistoryFile() error {
	content, err := os.ReadFile(historyFilePath)
	if err != nil {
		return err
	}

	if len(content) == 0 {
		log.Warn("History file is empty")
		return nil
	}

	return nil
}
