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
		logs.Error("Error during docker client initialization: ", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		option := os.Args[1]

		switch option {
		case "--images", "--image", "-i":
			err := img.ImageMode(ctx, cli)
			if err != nil {
				logs.Error("Error: ", err)
			}
			os.Exit(0)
		case "--build", "-build", "-b":
			err := img.BuildMode(ctx, cli)
			if err != nil {
				logs.Error("Error: ", err)
			}
			os.Exit(0)
		case "--all", "-a":
			showAllContainers = true
			con.ContainerMode(ctx, cli, showAllContainers)
		case "--help", "-h":
			printHelpManual()
			os.Exit(0)
		case "--version", "-v":
			printAsciiArt()
			fmt.Println("dk version 0.2.2")
			os.Exit(0)
		default:
			logs.WarnMsg(fmt.Sprintf("Unknown option: %s", option))
			printHelpManual()
			os.Exit(0)
		}
	}

	con.ContainerMode(ctx, cli, showAllContainers)
}

func printHelpManual() {
	fmt.Println("Usage: dk [options]")
	fmt.Printf("  %-20s %s\n", "dk", "Run the program")
	fmt.Printf("  %-20s %s\n", "dk [--all | -a]", "Run the program including all containers")
	fmt.Printf("  %-20s %s\n", "dk [--images | -i]", "Run image mode")
	fmt.Printf("  %-20s %s\n", "dk [--build | -b]", "Build a new image from a Dockerfile in the current directory")
	fmt.Printf("  %-20s %s\n", "dk [--help | -h]", "Show this help message")
}

func printAsciiArt() {
	ascii, err := os.ReadFile("./ressources/ascii.txt")
    if err != nil {
        fmt.Println("Erreur lors de la lecture du fichier:", err)
        return
    }

    fmt.Println(string(ascii))
}
