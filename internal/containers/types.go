package containers

import (
	"github.com/abroudoux/dk/internal/types"
	ct "github.com/docker/docker/api/types/container"
)

type Container = types.Container

type containerChoice struct {
	containers        []Container
	cursor            int
	containerSelected Container
}

type containerAction int
type actionChoice struct {
	actions           []containerAction
	cursor            int
	actionSelected    containerAction
	containerSelected Container
}

type containerStatus int
type statusChoice struct {
	statuses          []containerStatus
	cursor            int
	statusSelected    containerStatus
	containerSelected Container
}

type StopOptions = ct.StopOptions
