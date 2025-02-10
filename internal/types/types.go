package types

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	imagetypes "github.com/docker/docker/api/types/image"
)

type Container = types.Container
type Image = imagetypes.Summary
type ContainerListOptions = container.ListOptions
