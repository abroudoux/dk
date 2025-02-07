package images

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/abroudoux/dk/internal/logs"
	"github.com/abroudoux/dk/internal/utils"
	"github.com/charmbracelet/huh"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func runImage(image Image, ctx context.Context, cli *client.Client) error {
    var (
        containerName    string
        ports            string
        env              bool
        envs             []string
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

            huh.NewConfirm().
                Title("Do you want to add environment variables?").
                Negative("No").
                Affirmative("Yes").
                Value(&env),
        ),
    )

    err := form.Run()
    if err != nil {
        return fmt.Errorf("Failed to run form: %v", err)
    }

    if ports != "" && !checkPortInput(ports) {
        return fmt.Errorf("Invalid port mapping")
    }

    if env {
        getEnvs(&envs)
    }

    config := &containertypes.Config{
        Image: image.ID,
        Tty:   true,
        Env:   envs,
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
		return fmt.Errorf("Failed to create container %s: %v", containerName, err)
	}

    if err := cli.ContainerStart(ctx, resp.ID, containertypes.StartOptions{}); err != nil {
        return fmt.Errorf("Failed to start container %s: %v", containerName, err)
    }

	logs.InfoMsg(fmt.Sprintf("Container %s started", containerName))
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

func getEnv() string {
    var key string;
    var value string;
    huh.NewInput().Title("Key").Prompt("? ").Value(&key).Run()
    huh.NewInput().Title("Value").Prompt("? ").Value(&value).Run()
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
    logs.InfoMsg(fmt.Sprintf("Environment variable saved: %s", newEnv))

    addNewEnv := utils.GetConfirmation("Do you want to add another environment variable?")
    if addNewEnv {
        getEnvs(envs)
        return
    }
    return
}