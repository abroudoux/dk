package history

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
)

func InitHistory() error {
	if !fileAlreadyExists(historyFilePath) {
		err := createHistoryFile(historyFilePath)
		if err != nil {
			return err
		}
	}
	return nil
}

func fileAlreadyExists(filePath string) bool {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func createHistoryFile(filePath string) error {
	_, err := os.Create(filePath)
	if err != nil {
		return err
	}

	log.Info("History file created at:", filePath)
	return nil
}
