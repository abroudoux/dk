package forms

import (
	"fmt"

	"github.com/abroudoux/dk/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

type Status int

const (
	StatusExit Status = iota
	StatusPause
	StatusRestart
	StatusStop
)

func (s Status) String() string {
	return [...]string{"Exit", "Pause", "Restart", "Stop"}[s]
}

type statusChoice struct {
	statuses []Status
	cursor int
	selectedStatus Status
	selectedContainer Container
}

func initialStatusModel(container Container) statusChoice {
	statuses := []Status{
		StatusExit,
		StatusPause,
		StatusRestart,
		StatusStop,
	}

	return statusChoice{
		statuses: statuses,
		cursor: len(statuses) - 1,
		selectedStatus: StatusPause,
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
	s += fmt.Sprintf("Container: %s\n\n", menu.selectedContainer.Names[0])

	for i, status := range menu.statuses {
		cursor := " "
		cursor = ui.RenderCursor(menu.cursor == i)
		s += fmt.Sprintf("%s %s\n", cursor, ui.RenderSelected(status.String(), menu.cursor == i))

	}
	return s
}

func ChooseStatus(container Container) (Status, error) {
	p := tea.NewProgram(initialStatusModel(container))
	m, err := p.Run()
	if err != nil {
		return StatusExit, err
	}

	status := m.(statusChoice).selectedStatus
	return status, nil
}