package images

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
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
    )

    form := huh.NewForm(
        huh.NewGroup(
            huh.NewInput().
                Title("Container name (--name) (Optional)").
                Value(&containerName),

            huh.NewInput().
                Title("Port mapping (-p)").
                Placeholder("host:container").
                Value(&ports),
        ),
    )

    err := form.Run()
    if err != nil {
        return err
    }

    if ports != "" && !checkPortInput(ports) {
        return nil
    }

    config := &containertypes.Config{
        Image: image.ID,
        Tty:   true,
    }

    hostConfig := &containertypes.HostConfig{
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

func checkPortInput(port string) bool {
    portRegex := regexp.MustCompile(`^(\d+s):(\d+)$`)
    if !portRegex.MatchString(port) {
        logs.ErrorMsg("Invalid port mapping. Please use the following format: host:container")
        return false
    }

    portParts := strings.Split(port, ":")
    _, errHost := strconv.Atoi(portParts[0])
    _, errContainer := strconv.Atoi(portParts[1])
    if errHost != nil || errContainer != nil {
        logs.ErrorMsg("Invalid port numbers. Please use integers for both host and container ports")
        return false
    }

    return true
}