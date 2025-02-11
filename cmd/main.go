package main

import (
	"fmt"
	"os"

	con "github.com/abroudoux/dk/internal/containers"
	"github.com/abroudoux/dk/internal/docker"
	img "github.com/abroudoux/dk/internal/images"
	"github.com/abroudoux/dk/internal/logs"
	"github.com/abroudoux/dk/internal/utils"
)

func main() {
	var showAllContainers bool = false

	ctx := utils.GetContext()
	cli, err := docker.GetCli()
	if err != nil {
		logs.Error("Error during docker Client initialization: ", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		option := os.Args[1]

		switch os.Args[1] {
		case "--all", "-a":
			showAllContainers = true
			err := con.ContainerMode(ctx, cli, showAllContainers)
			if err != nil {
				logs.Error("Error: ", err)
				os.Exit(1)
			}
		case "--images", "--image", "-i":
			err := img.ImageMode(ctx, cli)
			if err != nil {
				logs.Error("Error: ", err)
				os.Exit(1)
			}
			os.Exit(0)
		case "--build", "-build", "-b":
			err := img.BuildMode(ctx, cli)
			if err != nil {
				logs.Error("Error: ", err)
			}
			os.Exit(0)
		case "--help", "-h":
			utils.PrintHelpManual()
			os.Exit(0)
		case "--version", "-v":
			err := utils.PrintAsciiArt()
			if err != nil {
				logs.Error("Failed to print ASCII art: ", err)
			}
			fmt.Println("dk version 0.3.0")
			os.Exit(0)
		default:
			logs.WarnMsg(fmt.Sprintf("Unknown option: %s", option))
			utils.PrintHelpManual()
			os.Exit(0)
		}
	}

	con.ContainerMode(ctx, cli, showAllContainers)
}
