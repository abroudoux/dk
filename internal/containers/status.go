package containers

import (
	"fmt"

	"github.com/abroudoux/dk/internal/ui"
	"github.com/abroudoux/dk/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

type ContainerStatus int

const (
	ContainerStatusExit ContainerStatus = iota
	ContainerStatusPause
	ContainerStatusRestart
	ContainerStatusStop
)

func (s ContainerStatus) String() string {
	return [...]string{"Exit", "Pause", "Restart", "Stop"}[s]
}

type statusChoice struct {
	statuses []ContainerStatus
	cursor int
	selectedStatus ContainerStatus
	selectedContainer Container
}

func initialStatusModel(container Container) statusChoice {
	statuses := []ContainerStatus{
		ContainerStatusExit,
		ContainerStatusPause,
		ContainerStatusRestart,
		ContainerStatusStop,
	}

	return statusChoice{
		statuses: statuses,
		cursor: len(statuses) - 1,
		selectedStatus: ContainerStatusExit,
		selectedContainer: container,
	}
}

func (menu statusChoice) Init() tea.Cmd {
	return nil
}

func (menu statusChoice) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return menu, tea.Quit
		case "up":
			menu.cursor--
			if menu.cursor < 0 {
				menu.cursor = len(menu.statuses) - 1
			}
		case "down":
			menu.cursor++
			if menu.cursor >= len(menu.statuses) {
				menu.cursor = 0
			}
		case "enter":
			menu.selectedStatus = menu.statuses[menu.cursor]
			return menu, tea.Quit
		}
	}

	return menu, nil
}

func (menu statusChoice) View() string {
	s := "\033[H\033[2J"
	s += fmt.Sprintf("Container: %s\n\n", utils.RenderContainerName(menu.selectedContainer))

	for i, status := range menu.statuses {
		cursor := " "
		cursor = ui.RenderCursor(menu.cursor == i)
		s += fmt.Sprintf("%s %s\n", cursor, ui.RenderSelected(status.String(), menu.cursor == i))

	}
	return s
}

func SelectStatus(container Container) (ContainerStatus, error) {
	p := tea.NewProgram(initialStatusModel(container))
	m, err := p.Run()
	if err != nil {
		return ContainerStatusExit, err
	}

	status := m.(statusChoice).selectedStatus
	return status, nil
}