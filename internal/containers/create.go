package containers

import t "github.com/abroudoux/dk/internal/types"

func CreateContainer(ctx t.Context, cli t.Client, config *ContainerConfig, hostConfig *ContainerHostConfig, containerName string) (containerId string, err error) {
	response, err := cli.ContainerCreate(ctx, config, hostConfig, nil, nil, containerName)
	if err != nil {
		return "", nil
	}

	return response.ID, nil
}
