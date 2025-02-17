package images

import (
	"fmt"
	"strings"

	containers "github.com/abroudoux/dk/internal/containers"
	"github.com/abroudoux/dk/internal/logs"
	t "github.com/abroudoux/dk/internal/types"
	"github.com/abroudoux/dk/internal/ui"
	"github.com/abroudoux/dk/internal/utils"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/docker/go-connections/nat"
)

func runForm(cmd *commandRun, image image) error {
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

			huh.NewInput().
				Title("Network (--network) (Leave empty if you don't want to specifies it)").
				Placeholder("Network name").
				Value(&cmd.network),

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
		return err
	}

	return nil
}

func runImage(image image, ctx t.Context, cli t.Client) error {
	var cmd commandRun

	utils.CleanView()
	err := runForm(&cmd, image)

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
		err = getHostPort(image, &cmd.ports, ctx, cli)
	}

	hostConfig.PortBindings = nat.PortMap{
		nat.Port(strings.Split(cmd.ports, ":")[1] + "/tcp"): []nat.PortBinding{
			{HostIP: "0.0.0.0", HostPort: strings.Split(cmd.ports, ":")[0]},
		},
	}

	if cmd.network != "" {
		hostConfig.NetworkMode = containers.ContainerNetworkMode(cmd.network)
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

func getEnv() string {
	var key string
	var value string
	huh.NewInput().Title("Key").Prompt("? ").Value(&key).Run()
	huh.NewInput().Title("Value").Prompt("? ").Value(&value).Run()

	if key == "" || value == "" {
		logs.WarnMsg("Key or value can't be empty")
		return getEnv()
	}

	return key + "=" + value
}

func isEnvAlreadySaved(newEnv string, envs *[]string) bool {
	for _, env := range *envs {
		if env == newEnv {
			return true
		}
	}
	return false
}

func getEnvs(envs *[]string) {
	newEnv := getEnv()
	envAlreadySaved := isEnvAlreadySaved(newEnv, envs)
	if envAlreadySaved {
		logs.WarnMsg("Environment variable already saved")
		getEnvs(envs)
		return
	}

	*envs = append(*envs, newEnv)
	log.Info(fmt.Sprintf("Environment variable saved: %s", newEnv))

	addNewEnv := utils.GetConfirmation("Do you want to add another environment variable?")
	if addNewEnv {
		getEnvs(envs)
		return
	}
	return
}

func getHostPort(image image, ports *string, ctx t.Context, cli t.Client) error {
	imageInspected, _, err := cli.ImageInspectWithRaw(ctx, image.ID)
	if err != nil {
		return err
	}

	var exposedPort string
	for port := range imageInspected.Config.ExposedPorts {
		exposedPort = string(port)
		break
	}

	if exposedPort == "" {
		logs.WarnMsg("No exposed ports found in the image")
		return nil
	}

	portNumber := strings.Split(exposedPort, "/")[0]
	*ports = fmt.Sprintf("%s:%s", portNumber, portNumber)
	return nil
}
