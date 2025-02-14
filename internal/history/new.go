package history

import t "github.com/abroudoux/dk/internal/types"

func NewImageCommand(cmdType CommandImageType, image t.Image) *CommandHistory {
	return &CommandHistory{
		Type:  CommandType(CommandTypeImage),
		Image: &image,
	}
}
