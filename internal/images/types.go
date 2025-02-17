package images

import (
	dkTypes "github.com/abroudoux/dk/internal/types"
	types "github.com/docker/docker/api/types"
	it "github.com/docker/docker/api/types/image"
)

type image = dkTypes.Image

type commandRun struct {
	containerName   string
	ports           string
	env             bool
	envs            []string
	removeContainer bool
	network         string
}

type imageAction int

type actionChoice struct {
	actions        []imageAction
	cursor         int
	actionSelected imageAction
	imageSelected  image
}

type imageBuildOptions = types.ImageBuildOptions
type imageListOptions = it.ListOptions
type imageRemoveOptions = it.RemoveOptions
