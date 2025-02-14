package history

import (
	"encoding/json"
	"time"

	t "github.com/abroudoux/dk/internal/types"
)

func SaveCommand(cmd *CommandHistory) error {}

func createJsonObj(cmd *CommandHistory) (string, error) {
	type JsonCommand struct {
		Context   string    `json:"context"`
		Type      string    `json:"type"`
		ImageInfo t.Image   `json:"image"`
		Timestamp time.Time `json:"timestamp"`
	}

	type ImageInfo struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Tag  string `json:"tag"`
	}

	var context, cmdType string
	var imageInfo t.Image

	switch cmd.Type {
	case CommandTypeImage:
		context = "image"
		switch cmd.Type {
		case CommandImageTypeRun:
			cmdType = "run"
		case CommandImageTypeBuild:
			cmdType = "build"
		case CommandImageTypeDelete:
			cmdType = "delete"
		}
		imageInfo = t.Image{
			ID:   cmd.Image.ID,
			Name: cmd.Image.Name,
			Tag:  cmd.Image.Tag,
		}
	case CommandTypeContainer:
		context = "container"
	}

	jsonCmd := JsonCommand{
		Context:   context,
		Type:      cmdType,
		ImageInfo: imageInfo,
		Timestamp: cmd.Timestamp,
	}

	jsonBytes, err := json.Marshal(jsonCmd)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}
