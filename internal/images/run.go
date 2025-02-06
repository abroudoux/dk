package images

import (
	"context"
	"fmt"
	"strings"

	"github.com/abroudoux/dk/internal/logs"
	"github.com/charmbracelet/huh"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func runImage(image Image, ctx context.Context, cli *client.Client) error {
    var (
        containerName    string
        ports            string
        // removeContainer  bool
    )

    form := huh.NewForm(
        huh.NewGroup(
            huh.NewInput().
                Title("Container name (--name)").
                Value(&containerName),

            huh.NewInput().
                Title("Port mapping (-p)").
                Placeholder("host:container").
                Value(&ports),

            // huh.NewConfirm().
            //     Title("Remove container when it exits (--rm)?").
            //     Value(&removeContainer),
        ),
    )

    err := form.Run()
    if err != nil {
        return err
    }

    config := &containertypes.Config{
        Image: image.ID,
        Tty:   true,
    }

    hostConfig := &containertypes.HostConfig{
        // AutoRemove: removeContainer,
        AutoRemove: false,
    }

    if ports != "" {
        hostConfig.PortBindings = nat.PortMap{
            nat.Port(strings.Split(ports, ":")[1] + "/tcp"): []nat.PortBinding{
                {HostIP: "0.0.0.0", HostPort: strings.Split(ports, ":")[0]},
            },
        }
    }

    resp, err := cli.ContainerCreate(ctx, config, hostConfig, nil, nil, containerName)
	if err != nil {
		return err
	}

    if err := cli.ContainerStart(ctx, resp.ID, containertypes.StartOptions{}); err != nil {
        return err
    }

	logs.InfoMsg(fmt.Sprintf("Container %s started", containerName))
    return nil
}