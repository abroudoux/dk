package utils

import (
	"strings"

	"github.com/abroudoux/dk/internal/ui"
	"github.com/docker/docker/api/types"
)

type Container = types.Container

func RenderContainerName(container Container) string {
	containerName := strings.Join(container.Names, "")
	containerNameWithoutPreffix := strings.Trim(containerName, "/")
	return ui.RenderContainer(containerNameWithoutPreffix)
}