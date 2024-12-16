package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
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

	containers, err := getContainers()
	if err != nil {
		println("Error getting containers")
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		flagMode(containers)
		os.Exit(0)
	}

	containerSelected, err := chooseContainer(containers)
	if err != nil {
		println("Error choosing container", err)
		os.Exit(1)
	}

	println("Container selected: ", containerSelected)
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

func flagMode(containers []string) {
	flag := os.Args[1]

	switch flag {
	case "--run", "-r":
		containerSelected, err := chooseContainer(containers)
		if err != nil {
			println("Error choosing container")
			os.Exit(1)
		}

		println("Container selected: ", containerSelected)
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
	cmd := exec.Command("docker", "container", "ls", "-a")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	var containers []string
	for _, line := range lines {
		if line != "" {
			containers = append(containers, line)
		}
	}

	return containers, nil
}

type containerChoice struct {
	containers []string
	cursor    int
	selectedContainer string
}

func initialContainerModel(containers []string) containerChoice {
	return containerChoice{
		containers: containers,
		cursor:    len(containers) - 1,
		selectedContainer: "",
	}
}

func (menu containerChoice) Init() tea.Cmd {
	return nil
}

func (menu containerChoice) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return menu, tea.Quit
		case "up":
			menu.cursor--
			if menu.cursor < 0 {
				menu.cursor = len(menu.containers) - 1
			}
		case "down":
			menu.cursor++
			if menu.cursor >= len(menu.containers) {
				menu.cursor = 0
			}
		case "enter":
			menu.selectedContainer = menu.containers[menu.cursor]
			return menu, tea.Quit
		}
	}

	return menu, nil
}

func (menu containerChoice) View() string {
	s := "\033[H\033[2J"
    s += "Choose a container:\n\n"

	for i, container := range menu.containers {
        cursor := " "

        if menu.cursor == i {
            // cursor = renderCursor()
            cursor = ">"
            s += fmt.Sprintf("%s %s\n", cursor, container)
        } else {
            s += fmt.Sprintf("%s %s\n", cursor, container)
        }
    }

    return s
}

func chooseContainer(containers []string) (string, error) {
	containersMenu := tea.NewProgram(initialContainerModel(containers))
	finalModel, err := containersMenu.Run()
	if err != nil {
		return "", err
	}

	containerMenu := finalModel.(containerChoice)
	return containerMenu.selectedContainer, nil
}