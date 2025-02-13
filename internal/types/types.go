package types

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	imagetypes "github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

type Container = types.Container
type Image = imagetypes.Summary
type ContainerListOptions = container.ListOptions

type Client = *client.Client
type Context = context.Context
