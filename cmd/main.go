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
	cli, err := docker.GetClient()
	if err != nil {
		utils.PrintErrorAndExit(err)
	}

	if len(os.Args) > 1 {
		option := os.Args[1]

		switch option {
		case "--all", "-a", "all":
			showAllContainers = true
		case "--images", "--image", "images", "image", "-i":
			err := img.ImageMode(ctx, cli)
			if err != nil {
				utils.PrintErrorAndExit(err)
			}
			os.Exit(0)
		case "--build", "-build", "build", "-b":
			err := img.BuildMode(ctx, cli)
			if err != nil {
				utils.PrintErrorAndExit(err)
			}
			os.Exit(0)
		case "--help", "-h":
			utils.PrintHelpManual()
			os.Exit(0)
		case "--version", "-v":
			utils.PrintAsciiArt()
			utils.PrintVersion()
			os.Exit(0)
		default:
			logs.WarnMsg(fmt.Sprintf("Unknown option: %s", option))
			utils.PrintHelpManual()
			os.Exit(0)
		}
	}

	con.ContainerMode(ctx, cli, showAllContainers)
	os.Exit(0)
}
