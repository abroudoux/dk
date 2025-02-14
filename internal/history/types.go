package history

import (
	t "github.com/abroudoux/dk/internal/types"
)

type CommandType int
type CommandContainerType int
type CommandContainerStatusType int
type CommandImageType int

const (
	CommandTypeContainer CommandType = iota
	CommandTypeImage
)

const (
	CommandContainerTypeCopyContainerId CommandContainerType = iota
	CommandContainerTypeDelete
	CommandContainerTypeLogs
	CommandContainerTypeStatus
)

const (
	CommandContainerStatusTypePause CommandContainerStatusType = iota
	CommandContainerStatusTypeRestart
	CommandContainerStatusTypeStop
)

const (
	CommandImageTypeRun CommandImageType = iota
	CommandImageTypeBuild
	CommandImageTypeDelete
)

type CommandHistory struct {
	Type      CommandType
	Image     *t.Image
	Container *t.Container
}
