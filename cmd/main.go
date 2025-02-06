package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	con "github.com/abroudoux/dk/internal/containers"
	"github.com/abroudoux/dk/internal/docker"
	img "github.com/abroudoux/dk/internal/images"
	"github.com/abroudoux/dk/internal/logs"
	"github.com/docker/docker/client"
)

func main() {
	var showAllContainers bool = false

	ctx, cli, err := docker.GetCtxCli()
	if err != nil {
		logs.Error("Error during docker client initialization: ", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		option := os.Args[1]

		switch option {
		case "--images", "--image", "-i":
			err := imageMode(ctx, cli)
			if err != nil {
				logs.Error("Error: ", err)
			}
			os.Exit(0)
		case "--all", "-a":
			showAllContainers = true
			containerMode(ctx, cli, showAllContainers)
		case "--help", "-h":
			PrintHelpManual()
			os.Exit(0)
		case "--version", "-v":
			fmt.Println("dk version 0.0.1")
			os.Exit(0)
		default:
			logs.WarnMsg(fmt.Sprintf("Unknown option: %s", option))
			PrintHelpManual()
			os.Exit(0)
		}
	}

	containerMode(ctx, cli, showAllContainers)
}

func containerMode(ctx context.Context, cli *client.Client, showAllContainers bool) {
	containers, err := con.GetContainers(ctx, cli, showAllContainers)
	if err != nil {
		logs.Error("Error during containers recuperation: ", err)
		os.Exit(1)
	}

	if len(containers) == 0 {
		logs.WarnMsg("No containers found")
		os.Exit(0)
	}

	sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		containerSelected, err := con.SelectContainer(containers)
		if err != nil {
			logs.Error("Error during container selection: ", err)
			os.Exit(1)
		}

		if containerSelected.ID == "" {
			logs.InfoMsg("No container selected. Exiting program.")
			os.Exit(0)
		}

		action, err := con.SelectAction(containerSelected)
		if err != nil {
			logs.Error("Error during action selection: ", err)
			os.Exit(1)
		}

		err = con.DoContainerAction(ctx, cli, containerSelected, action)
		if err != nil {
			logs.Error("Error during action execution: ", err)
			os.Exit(1)
		}

		os.Exit(0)
    }()

	<-sigChan
    fmt.Println("\nProgram interrupted. Exiting...")
    os.Exit(0)
}

func imageMode(ctx context.Context, cli *client.Client) error {
	images, err := img.GetImages(ctx, cli, true)
	if err != nil {
		logs.Error("Error during images recuperation: ", err)
		os.Exit(1)
	}

	imageSelected, err := img.SelectImage(images)
	if err != nil {
		return err
	}

	action, err := img.SelectAction(imageSelected)
	if err != nil {
		return err
	}

	err = img.DoImageAction(ctx, cli, imageSelected, action)
	if err != nil {
		return err
	}

	return nil
}


func PrintHelpManual() {
	fmt.Println("Usage: dk [options]")
	fmt.Printf("  %-20s %s\n", "dk", "Run the program")
	fmt.Printf("  %-20s %s\n", "dk [--help | -h]", "Show this help message")
}
