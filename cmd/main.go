package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/abroudoux/dk/internal/container"
	"github.com/abroudoux/dk/internal/docker"
	"github.com/abroudoux/dk/internal/forms"
	"github.com/abroudoux/dk/internal/logs"
)

func main() {
	var allContainers bool

	if len(os.Args) > 1 {
		option := os.Args[1]

		switch option {
		case "--all", "-a":
			allContainers = true
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

	ctx, cli, err := docker.GetCtxCli()
	if err != nil {
		logs.Error("Error during docker client initialization: ", err)
		os.Exit(1)
	}

	containers, err := container.GetContainers(ctx, cli, allContainers)
	if err != nil {
		logs.Error("Error during containers recuperation: ", err)
		os.Exit(1)
	}

	sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		containerSelected, err := forms.ChooseContainer(containers)
		if err != nil {
			logs.Error("Error during container selection: ", err)
		}

		println(containerSelected.ID)

		action, err := forms.ChooseAction(containerSelected)
		if err != nil {
			logs.Error("Error during action selection: ", err)
		}

		fmt.Printf("%s", action)
		os.Exit(0)
    }()

	<-sigChan
    fmt.Println("\nProgram interrupted. Exiting...")
    os.Exit(0)
}


func PrintHelpManual() {
	fmt.Println("Usage: dk [options]")
	fmt.Printf("  %-20s %s\n", "dk", "Run the program")
	fmt.Printf("  %-20s %s\n", "dk [--help | -h]", "Show this help message")
}
