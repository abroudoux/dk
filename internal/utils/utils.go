package utils

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/abroudoux/dk/internal/logs"
	"github.com/abroudoux/dk/internal/types"
	"github.com/abroudoux/dk/internal/ui"
	"github.com/charmbracelet/huh"
)

type Container = types.Container
type Image = types.Image

func GetContext() context.Context {
	ctx := context.Background()
	return ctx
}

func RenderContainerName(container Container) string {
	containerName := strings.Join(container.Names, "")
	containerNameWithoutPreffix := strings.Trim(containerName, "/")
	return ui.RenderElementSelected(containerNameWithoutPreffix)
}

func RenderImageName(image Image) string {
	if len(image.RepoTags) == 0 {
		return ui.RenderElementSelected("<none>")
	} else {
		imageName := strings.Join(image.RepoTags, "")
		return ui.RenderElementSelected(imageName)
	}
}

func FormatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%dB", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%cB", float64(size)/float64(div), "KMGTPE"[exp])
}

func GetConfirmation(message string) bool {
	var confirm bool
	huh.NewConfirm().Title(message).Negative("No").Affirmative("Yes!").Value(&confirm).Run()
	return confirm
}

func CleanView() {
	println("\033[H\033[2J")
}

func PrintHelpManual() {
	commands := []string{"dk", "dk [--all | -a]", "dk [--images | -i]", "dk [--build | -b]", "dk [--help | -h]"}
	descriptions := []string{"Run the program", "Run the program including all containers", "Run image mode", "Build a new image from a Dockerfile in the current directory", "Show this help message"}

	fmt.Println("Usage: dk [options]")
	for i, cmd := range commands {
		fmt.Printf("  %-20s %s\n", cmd, descriptions[i])
	}
}

func PrintAsciiArt() error {
	ascii, err := os.ReadFile("./ressources/ascii.txt")
	if err != nil {
		return err
	}

	fmt.Println(string(ascii))
	return nil
}

func PrintErrorAndExit(err error) {
	logs.Error("Error", err)
	os.Exit(1)
}
