package containers

import (
	"fmt"

	"github.com/abroudoux/dk/internal/ui"
	"github.com/abroudoux/dk/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

type ContainerAction int

const (
	ContainerActionExit ContainerAction = iota
	ContainerActionCopyContainerID
	ContainerActionDelete
	ContainerActionLogs
	ContainerActionsStatus
)

func (a ContainerAction) String() string {
	return [...]string{"Exit", "Copy Container ID", "Delete", "Logs", "Status"}[a]
}

type actionChoice struct {
	actions []ContainerAction
	cursor int
	selectedAction ContainerAction
	selectedContainer Container
}

func initialActionModel(container Container) actionChoice {
	actions := []ContainerAction{
		ContainerActionExit,
		ContainerActionCopyContainerID,
		ContainerActionDelete,
		ContainerActionLogs,
		ContainerActionsStatus,
	}

	return actionChoice{
		actions: actions,
		cursor: len(actions) - 1,
		selectedAction: ContainerActionExit,
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
	s := "\033[H\033[2J\n"
	s += fmt.Sprintf("Container: %s\n\n", utils.RenderContainerName(menu.selectedContainer))

	for i, action := range menu.actions {
		cursor := " "
		cursor = ui.RenderCursor(menu.cursor == i)
		s += fmt.Sprintf("%s %s\n", cursor, ui.RenderLineSelected(action.String(), menu.cursor == i))
	}

	return s
}

func selectAction(container Container) (ContainerAction, error) {
	p := tea.NewProgram(initialActionModel(container))
	m, err := p.Run()
	if err != nil {
		return ContainerActionExit, err
	}

	action := m.(actionChoice).selectedAction
	return action, nil
}