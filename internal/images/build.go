package images

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/abroudoux/dk/internal/logs"
	"github.com/charmbracelet/huh"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

func buildImage(ctx context.Context, cli *client.Client) error {
    var imageName string

    if _, err := os.Stat("Dockerfile"); os.IsNotExist(err) {
        logs.ErrorMsg("Dockerfile not found in current directory")
        return fmt.Errorf("dockerfile not found", err)
    }

    form := huh.NewForm(
        huh.NewGroup(
            huh.NewInput().
                Title("Image name").
                Value(&imageName),
        ),
    )

    err := form.Run()
    if err != nil {
        return fmt.Errorf("failed to get image name: %v", err)
    }

    buildContext, err := archive.TarWithOptions(".", &archive.TarOptions{})
    if err != nil {
        return fmt.Errorf("failed to create build context: %v", err)
    }

    options := types.ImageBuildOptions{
        Tags:        []string{imageName},
        Dockerfile:  "Dockerfile",
        Remove:      true,
        ForceRemove: true,
    }

    resp, err := cli.ImageBuild(ctx, buildContext, options)
    if err != nil {
        return fmt.Errorf("failed to build image: %v", err)
    }
    defer resp.Body.Close()

    decoder := json.NewDecoder(resp.Body)
    for {
        var message struct {
            Stream string `json:"stream"`
            Status string `json:"status"`
            Error  string `json:"error"`
        }

        if err := decoder.Decode(&message); err != nil {
            if err == io.EOF {
                break
            }
            return err
        }

        if message.Error != "" {
            logs.ErrorMsg(message.Error)
        } else if message.Stream != "" {
            trimmedStream := strings.TrimSpace(message.Stream)
            if trimmedStream != "" {
                logs.InfoMsg(trimmedStream)
            }
        } else if message.Status != "" {
            logs.InfoMsg(message.Status)
        }
    }

    logs.InfoMsg(fmt.Sprintf("Image %s built successfully", imageName))

    err = PruneImages(ctx, cli)
    if err != nil {
        return err
    }

    logs.InfoMsg("Image built and intermediate images cleaned up")
    return nil
}