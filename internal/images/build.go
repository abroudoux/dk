package images

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/abroudoux/dk/internal/logs"
	t "github.com/abroudoux/dk/internal/types"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/docker/docker/pkg/archive"
)

func buildImage(ctx t.Context, cli t.Client) error {
	var (
		imageName    string
		filePath     string
		buildContext io.Reader
		options      imageBuildOptions
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Image name").
				Value(&imageName),

			huh.NewInput().
				Title("Path to Dockerfile (Optional)").
				Value(&filePath),
		),
	)

	err := form.Run()
	if err != nil {
		return fmt.Errorf("failed to get image name: %v", err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %v", err)
	}

	if filePath != "" {
		filePathAbs := filePath
		if !filepath.IsAbs(filePath) {
			filePathAbs = filepath.Join(pwd, filePath)
		}

		if _, err := os.Stat(filePathAbs); errors.Is(err, os.ErrNotExist) {
			logs.ErrorMsg("Dockerfile not found in specified path")
			return fmt.Errorf("dockerfile not found: %v", err)
		}

		buildContextDir := filepath.Dir(filePathAbs)

		buildContext, err = archive.TarWithOptions(buildContextDir, &archive.TarOptions{})
		if err != nil {
			return fmt.Errorf("failed to create build context: %v", err)
		}

		options = imageBuildOptions{
			Tags:        []string{imageName},
			Dockerfile:  filepath.Base(filePathAbs),
			Remove:      true,
			ForceRemove: true,
		}
	} else {
		if _, err := os.Stat("Dockerfile"); errors.Is(err, os.ErrNotExist) {
			logs.ErrorMsg("Dockerfile not found in current directory")
			return fmt.Errorf("dockerfile not found: %v", err)
		}

		buildContext, err = archive.TarWithOptions(".", &archive.TarOptions{})
		if err != nil {
			return fmt.Errorf("failed to create build context: %v", err)
		}

		options = imageBuildOptions{
			Tags:        []string{imageName},
			Dockerfile:  "Dockerfile",
			Remove:      true,
			ForceRemove: true,
		}
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
				log.Info(trimmedStream)
			}
		} else if message.Status != "" {
			log.Info(message.Status)
		}
	}

	log.Info(fmt.Sprintf("Image %s built successfully", imageName))

	err = prune(ctx, cli)
	if err != nil {
		return err
	}

	log.Info("Image built and intermediate images cleaned up")
	return nil
}
