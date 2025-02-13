package containers

import (
	"fmt"

	"github.com/abroudoux/dk/internal/ui"
	"github.com/abroudoux/dk/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	containerStatusExit containerStatus = iota
	containerStatusPause
	containerStatusRestart
	containerStatusStop
)

func (s containerStatus) String() string {
	return [...]string{
		"Exit",
		"Pause",
		"Restart",
		"Stop",
	}[s]
}

func initialStatusModel(container container) statusChoice {
	statuses := []containerStatus{
		containerStatusExit,
		containerStatusPause,
		containerStatusRestart,
		containerStatusStop,
	}

	return statusChoice{
		statuses:          statuses,
		cursor:            len(statuses) - 1,
		statusSelected:    containerStatusExit,
		containerSelected: container,
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
			menu.statusSelected = containerStatusExit
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
			menu.statusSelected = menu.statuses[menu.cursor]
			return menu, tea.Quit
		}
	}

	return menu, nil
}

func (menu statusChoice) View() string {
	s := "\033[H\033[2J\n"
	s += fmt.Sprintf("Container: %s\n\n", utils.RenderContainerName(menu.containerSelected))

	for i, status := range menu.statuses {
		cursor := " "
		cursor = ui.RenderCursor(menu.cursor == i)
		s += fmt.Sprintf("%s %s\n", cursor, ui.RenderLineSelected(status.String(), menu.cursor == i))

	}
	return s
}

func selectStatus(container container) (containerStatus, error) {
	p := tea.NewProgram(initialStatusModel(container))
	m, err := p.Run()
	if err != nil {
		return containerStatusExit, err
	}

	status := m.(statusChoice).statusSelected
	return status, nil
}
