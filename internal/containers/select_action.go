package containers

import (
	"fmt"

	"github.com/abroudoux/dk/internal/ui"
	"github.com/abroudoux/dk/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	containerActionExit containerAction = iota
	containerActionCopyContainerID
	containerActionDelete
	containerActionLogs
	containerActionsStatus
)

func (a containerAction) String() string {
	return [...]string{
		"Exit",
		"Copy Container ID",
		"Delete",
		"Logs",
		"Status",
	}[a]
}

func initialActionModel(container container) actionChoice {
	actions := []containerAction{
		containerActionExit,
		containerActionCopyContainerID,
		containerActionDelete,
		containerActionLogs,
		containerActionsStatus,
	}

	return actionChoice{
		actions:           actions,
		cursor:            len(actions) - 1,
		actionSelected:    containerActionExit,
		containerSelected: container,
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
			menu.actionSelected = containerActionExit
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
			menu.actionSelected = menu.actions[menu.cursor]
			return menu, tea.Quit
		}
	}

	return menu, nil
}

func (menu actionChoice) View() string {
	s := "\033[H\033[2J\n"
	s += fmt.Sprintf("Container: %s\n\n", utils.RenderContainerName(menu.containerSelected))

	for i, action := range menu.actions {
		cursor := " "
		cursor = ui.RenderCursor(menu.cursor == i)
		s += fmt.Sprintf("%s %s\n", cursor, ui.RenderLineSelected(action.String(), menu.cursor == i))
	}

	s += "\n"

	return s
}

func selectAction(container container) (containerAction, error) {
	p := tea.NewProgram(initialActionModel(container))
	m, err := p.Run()
	if err != nil {
		return containerActionExit, err
	}

	action := m.(actionChoice).actionSelected
	return action, nil
}
