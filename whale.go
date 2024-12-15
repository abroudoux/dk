package main

import (
	"os"
	"os/exec"
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