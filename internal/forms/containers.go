package forms

import (
	"fmt"
	"strings"

	"github.com/abroudoux/dk/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/docker/docker/api/types"
)

type Container = types.Container

type containerChoice struct {
	containers []Container
	cursor int
	selectedContainer Container
}

func initialContainerModel(containers []Container) containerChoice {
	return containerChoice{
		containers: containers,
		cursor: len(containers) - 1,
		selectedContainer: Container{},
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
		name, _ := strings.CutPrefix(container.Names[0], "/")
		imageName := container.Image
		publicPort := container.Ports[0].PublicPort
		privatePort := container.Ports[0].PrivatePort
		containerLine := fmt.Sprintf("%s => %s (%d:%d)", name, imageName, publicPort, privatePort)

        cursor := " "
		cursor = ui.RenderCursor(menu.cursor == i)
		s += fmt.Sprintf("%s %s\n", cursor, ui.RenderContainerSelected(containerLine, menu.cursor == i))
    }

    return s
}

func ChooseContainer(containers []Container) (Container, error) {
	containersMenu := tea.NewProgram(initialContainerModel(containers))
	model, err := containersMenu.Run()
	if err != nil {
		return Container{}, err
	}
	selectedContainer := model.(containerChoice).selectedContainer
	return selectedContainer, nil
}
