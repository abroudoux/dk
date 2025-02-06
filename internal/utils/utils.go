package utils

import (
	"fmt"
	"strings"

	"github.com/abroudoux/dk/internal/types"
	"github.com/abroudoux/dk/internal/ui"
)

type Container = types.Container
type Image = types.Image

func RenderContainerName(container Container) string {
	containerName := strings.Join(container.Names, "")
	containerNameWithoutPreffix := strings.Trim(containerName, "/")
	return ui.RenderContainer(containerNameWithoutPreffix)
}

func RenderImageName(image Image) string {
    imageName := strings.Join(image.RepoTags, "")
    return ui.RenderElementSelected(imageName)
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