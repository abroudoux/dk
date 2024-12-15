package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if !isDockerInstalled() {
		println("Docker is not installed")
		os.Exit(1)
	}

	if !isDockerRunning() {
		println("Docker is not running")
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		flagMode()
		os.Exit(0)
	}

	containers, err := getContainers()
	if err != nil {
		println("Error getting containers")
		os.Exit(1)
	}

	for _, container := range containers {
		fmt.Println(container)
	}
}

func isDockerInstalled() bool {
	cmd := exec.Command("docker", "-v")
	err := cmd.Run()
	return err == nil
}

func isDockerRunning() bool {
	cmd := exec.Command("docker", "container", "ls")
	err := cmd.Run()
	return err == nil
}

func flagMode() {
	flag := os.Args[1]

	switch flag {
	case "--help", "-h":
		printHelpManual()
	case "--version", "-v":
		fmt.Println("0.0.1")
	}
}

func printHelpManual() {
	fmt.Println("Usage: whale [options]")
	fmt.Printf("  %-20s %s\n", "whale [--help | -h]", "Show this help message")
}

func getContainers() ([]string, error) {
	cmd := exec.Command("docker", "container", "ls", "-a",)
	containers, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return strings.Fields(string(containers)), nil
}