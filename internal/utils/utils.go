package utils

import (
	"fmt"
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

func FormatSize(size int64) string {
    const unit = 1024
    if size < unit {
        return fmt.Sprintf("%dB", size)
    }
    div, exp := int64(unit), 0
    for n := size / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return fmt.Sprintf("%.1f%cB", float64(size)/float64(div), "KMGTPE"[exp])
}