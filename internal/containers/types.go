package containers

import (
	t "github.com/docker/docker/api/types"
	ct "github.com/docker/docker/api/types/container"
)

type container = t.Container

type containerChoice struct {
	containers        []container
	cursor            int
	containerSelected container
}

type containerAction int
type actionChoice struct {
	actions           []containerAction
	cursor            int
	actionSelected    containerAction
	containerSelected container
}

type containerStatus int
type statusChoice struct {
	statuses          []containerStatus
	cursor            int
	statusSelected    containerStatus
	containerSelected container
}

type containerRemoveOptions = ct.RemoveOptions
type stopOptions = ct.StopOptions
type containerListOptions = ct.ListOptions

type ContainerStartOptions = ct.StartOptions
type ContainerConfig = ct.Config
type ContainerHostConfig = ct.HostConfig
