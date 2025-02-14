package history

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
)

func InitHistory() error {
	if !fileAlreadyExists() {
		err := createHistoryFile()
		if err != nil {
			return err
		}
	}
	return nil
}

func fileAlreadyExists() bool {
	if _, err := os.Stat(historyFilePath); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func createHistoryFile() error {
	_, err := os.Create(historyFilePath)
	if err != nil {
		return err
	}

	log.Info("History file created at:", historyFilePath)
	return nil
}
