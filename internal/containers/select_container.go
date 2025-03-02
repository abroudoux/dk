package containers

import (
	"fmt"
	"strings"
	"time"

	"github.com/abroudoux/dk/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func initialContainerModel(containers []container) containerChoice {
	return containerChoice{
		containers:        containers,
		cursor:            len(containers) - 1,
		containerSelected: container{},
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
			menu.containerSelected = container{}
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
			menu.containerSelected = menu.containers[menu.cursor]
			return menu, tea.Quit
		}
	}

	return menu, nil
}

func (menu containerChoice) View() string {
	s := "\033[H\033[2J\n"
	s += "Choose a container:\n\n"

	for i, container := range menu.containers {
		var containerLine string
		var publicPort uint16
		var privatePort uint16

		name, _ := strings.CutPrefix(container.Names[0], "/")
		imageName := container.Image
		state := container.State
		created := time.Unix(container.Created, 0).Format("2006-01-02 15:04:05")

		if len(container.Ports) > 0 {
			publicPort = container.Ports[0].PublicPort
			privatePort = container.Ports[0].PrivatePort
			containerLine = fmt.Sprintf("%s => %s (%d:%d) [%s - %s]", name, imageName, publicPort, privatePort, state, created)
		} else {
			containerLine = fmt.Sprintf("%s => %s [%s - %s]", name, imageName, state, created)
		}

		cursor := " "
		cursor = ui.RenderCursor(menu.cursor == i)
		s += fmt.Sprintf("%s %s\n", cursor, ui.RenderLineSelected(containerLine, menu.cursor == i))
	}

	return s
}

func selectContainer(containers []container) (container, error) {
	p := tea.NewProgram(initialContainerModel(containers))
	m, err := p.Run()
	if err != nil {
		return container{}, err
	}
	container := m.(containerChoice).containerSelected
	return container, nil
}
