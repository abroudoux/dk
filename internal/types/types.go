package types

import (
	"context"

	"github.com/docker/docker/api/types"
	it "github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

type Image = it.Summary
type Container = types.Container

type Client = *client.Client
type Context = context.Context
type App struct {
	client Client
	ctx    Context
}
