package images

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/abroudoux/dk/internal/containers"
	"github.com/abroudoux/dk/internal/logs"
	t "github.com/abroudoux/dk/internal/types"
	"github.com/abroudoux/dk/internal/ui"
	"github.com/abroudoux/dk/internal/utils"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/docker/go-connections/nat"
)

func runImage(image image, ctx t.Context, cli t.Client) error {
	var cmd commandRun

	utils.CleanView()
	fmt.Printf("Image %s \n\n", utils.RenderImageName(image))

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Container name (--name) (Optionnal)").
				Value(&cmd.containerName),

			huh.NewInput().
				Title("Port mapping (-p) (Leave empty to use exposed port)").
				Placeholder("host:container").
				Value(&cmd.ports),

			huh.NewConfirm().
				Title("Do you want to add environment variables?").
				Negative("No").
				Affirmative("Yes").
				Value(&cmd.env),

			huh.NewConfirm().
				Title("Do you want to remove the container after stopping it?").
				Negative("No").
				Affirmative("Yes").
				Value(&cmd.removeContainer),
		),
	)

	err := form.Run()
	if err != nil {
		return fmt.Errorf("Failed to run form: %v", err)
	}

	if cmd.env {
		getEnvs(&cmd.envs)
	}

	config := &containers.ContainerConfig{
		Image: image.ID,
		Tty:   true,
		Env:   cmd.envs,
	}

	hostConfig := &containers.ContainerHostConfig{
		AutoRemove: cmd.removeContainer,
	}

	if cmd.ports == "" {
		imageInspect, _, err := cli.ImageInspectWithRaw(ctx, image.ID)
		if err != nil {
			return fmt.Errorf("Failed to inspect image: %v", err)
		}

		var exposedPort string
		for port := range imageInspect.Config.ExposedPorts {
			exposedPort = string(port)
			break
		}

		if exposedPort != "" {
			portNumber := strings.Split(exposedPort, "/")[0]
			cmd.ports = fmt.Sprintf("%s:%s", portNumber, portNumber)
		} else {
			logs.WarnMsg("No exposed ports found in the image")
		}
	}

	if !checkPortInput(cmd.ports) {
		return fmt.Errorf("Invalid port mapping")
	}

	hostConfig.PortBindings = nat.PortMap{
		nat.Port(strings.Split(cmd.ports, ":")[1] + "/tcp"): []nat.PortBinding{
			{HostIP: "0.0.0.0", HostPort: strings.Split(cmd.ports, ":")[0]},
		},
	}

	resp, err := cli.ContainerCreate(ctx, config, hostConfig, nil, nil, cmd.containerName)
	if err != nil {
		return fmt.Errorf("Failed to create container %s: %v", ui.RenderElementSelected(cmd.containerName), err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, containers.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("Failed to start container %s: %v", ui.RenderElementSelected(cmd.containerName), err)
	}

	log.Info(fmt.Sprintf("Container %s based on %s started", ui.RenderElementSelected(cmd.containerName), utils.RenderImageName(image)))
	return nil
}

func checkPortInput(port string) bool {
	portRegex := regexp.MustCompile(`^(\d+):(\d+)$`)
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
