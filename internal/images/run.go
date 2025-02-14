package images

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/abroudoux/dk/internal/containers"
	"github.com/abroudoux/dk/internal/history"
	"github.com/abroudoux/dk/internal/logs"
	t "github.com/abroudoux/dk/internal/types"
	"github.com/abroudoux/dk/internal/ui"
	"github.com/abroudoux/dk/internal/utils"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/docker/go-connections/nat"
)

func runImage(image image, ctx t.Context, cli t.Client) error {
	var (
		containerName   string
		ports           string
		env             bool
		envs            []string
		removeContainer bool
	)

	utils.CleanView()
	fmt.Printf("Image %s \n\n", utils.RenderImageName(image))

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Container name (--name) (Optionnal)").
				Value(&containerName),

			huh.NewInput().
				Title("Port mapping (-p) (Leave empty to use exposed port)").
				Placeholder("host:container").
				Value(&ports),

			huh.NewConfirm().
				Title("Do you want to add environment variables?").
				Negative("No").
				Affirmative("Yes").
				Value(&env),

			huh.NewConfirm().
				Title("Do you want to remove the container after stopping it?").
				Negative("No").
				Affirmative("Yes").
				Value(&removeContainer),
		),
	)

	err := form.Run()
	if err != nil {
		return fmt.Errorf("Failed to run form: %v", err)
	}

	if env {
		getEnvs(&envs)
	}

	config := &containers.ContainerConfig{
		Image: image.ID,
		Tty:   true,
		Env:   envs,
	}

	hostConfig := &containers.ContainerHostConfig{
		AutoRemove: removeContainer,
	}

	if ports == "" {
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
			ports = fmt.Sprintf("%s:%s", portNumber, portNumber)
		} else {
			logs.WarnMsg("No exposed ports found in the image")
		}
	}

	if !checkPortInput(ports) {
		return fmt.Errorf("Invalid port mapping")
	}

	hostConfig.PortBindings = nat.PortMap{
		nat.Port(strings.Split(ports, ":")[1] + "/tcp"): []nat.PortBinding{
			{HostIP: "0.0.0.0", HostPort: strings.Split(ports, ":")[0]},
		},
	}

	resp, err := cli.ContainerCreate(ctx, config, hostConfig, nil, nil, containerName)
	if err != nil {
		return fmt.Errorf("Failed to create container %s: %v", ui.RenderElementSelected(containerName), err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, containers.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("Failed to start container %s: %v", ui.RenderElementSelected(containerName), err)
	}

	log.Info(fmt.Sprintf("Container %s based on %s started", ui.RenderElementSelected(containerName), utils.RenderImageName(image)))

	cmd := history.NewImageCommand(history.CommandImageTypeRun, image)

	log.Info("New command added to history")
	println("Image Id:", cmd.Image.ID)

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
