package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

//go:embed config.json
var configFile string
var config Config

type Config struct {
	Ui struct {
		CursorColor string `json:"cursorColor"`
		BranchColor string `json:"branchColor"`
		ContainerSelectedColor string `json:"containerSelectedColor"`
		ActionSelectedColor string `json:"actionSelectedColor"`
	} `json:"Ui"`
}

func main() {
	err := loadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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

	for _, container := range containers {
		println(container)
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

	actionSelected, err := chooseAction(containerSelected)
	if err != nil {
		println("Error choosing action", err)
		os.Exit(1)
	}

	err = doAction(actionSelected, containerSelected)
	if err != nil {
		println("Error doing action", err)
		os.Exit(1)
	}
}

func loadConfig() error {
	err := json.Unmarshal([]byte(configFile), &config)
	if err != nil {
		return fmt.Errorf("error parsing config file: %v", err)
	}

	return nil
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

		actionSelected, err := chooseAction(containerSelected)
		if err != nil {
			println("Error choosing action")
			os.Exit(1)
		}

		println("Action selected: ", actionSelected)
	case "--help", "-h":
		printHelpManual()
	case "--version", "-v":
		fmt.Println("0.0.1")
	}
}

func printHelpManual() {
	fmt.Println("Usage: whale [options]")
	fmt.Printf("  %-20s %s\n", "whale [--run | -r]", "Run the program")
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
	for i, line := range lines {
		if i > 0 && line != "" {
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
            cursor = renderCursor()
            s += fmt.Sprintf("%s %s\n", cursor, renderContainerSelected(container, true))
        } else {
            s += fmt.Sprintf("%s %s\n", cursor, renderContainerSelected(container, false))
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

func renderCursor() string {
	render := fmt.Sprintf("\033[%sm>\033[0m", config.Ui.CursorColor)
	return render
}

func renderContainerSelected(container string, isSelected bool) string {
    if isSelected {
        return fmt.Sprintf("\033[%sm%s\033[0m", config.Ui.ContainerSelectedColor, container)
    }
    return container
}

type actionChoice struct {
	actions []string
	cursor int
	selectedAction string
	selectedContainer string
}

func initialActionModel(container string) actionChoice {
	actions := []string{
		"Exit",
		"Copy container ID",
	}

	return actionChoice{
		actions: actions,
		cursor: len(actions) - 1,
		selectedAction: "",
		selectedContainer: container,
	}
}

func (menu actionChoice) Init() tea.Cmd {
	return nil
}

func (menu actionChoice) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return menu, tea.Quit
		case "up":
			menu.cursor--
			if menu.cursor < 0 {
				menu.cursor = len(menu.actions) - 1
			}
		case "down":
			menu.cursor++
			if menu.cursor >= len(menu.actions) {
				menu.cursor = 0
			}
		case "enter":
			menu.selectedAction = menu.actions[menu.cursor]
			return menu, tea.Quit
		}
	}

	return menu, nil
}

func (menu actionChoice) View() string {
	s := "\033[H\033[2J"
	s += fmt.Sprintf("Container: %s\n\n", strings.Split(menu.selectedContainer, " ")[0])

	for i, action := range menu.actions {
		cursor := " "

		if menu.cursor == i {
			cursor = renderCursor()
			s += fmt.Sprintf("%s %s\n", cursor, renderActionSelected(action, true))
		} else {
			s += fmt.Sprintf("%s %s\n", cursor, renderActionSelected(action, false))
		}
	}

	return s
}

func renderActionSelected(action string, isSelected bool) string {
    if isSelected {
        return fmt.Sprintf("\033[%sm%s\033[0m", config.Ui.ActionSelectedColor, action)
    }
    return action
}

func chooseAction(container string) (string, error) {
	actionsMenu := tea.NewProgram(initialActionModel(container))
	finalModel, err := actionsMenu.Run()
	if err != nil {
		return "", err
	}

	actionMenu := finalModel.(actionChoice)
	return actionMenu.selectedAction, nil
}

func doAction(action string, container string) error {
	switch action {
	case "Exit":
		os.Exit(0)
	case "Copy container ID":
		containerID := strings.Split(container, " ")[0]
		err := copyContainerId(containerID)
		if err != nil {
			println(err)
			os.Exit(1)
		}
	}

	return nil
}

func copyContainerId(container string) error {
	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(container)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error copying container ID: %v", err)
	}

	println("Container ID copied to clipboard")
	return nil
}