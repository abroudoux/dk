package forms

import (
	"fmt"

	"github.com/abroudoux/dk/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

type actionChoice struct {
	actions []string
	cursor int
	selectedAction string
	selectedContainer Container
}

func initialActionModel(container Container) actionChoice {
	actions := []string{
		"Exit",
		"Copy Container ID",
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
	s += fmt.Sprintf("Container: %s\n\n", menu.selectedContainer.Names[0])

	for i, action := range menu.actions {
		cursor := " "
		cursor = ui.RenderCursor(menu.cursor == i)
		s += fmt.Sprintf("%s %s\n", cursor, ui.RenderActionSelected(action, menu.cursor == i))
	}

	return s
}

func ChooseAction(container Container) (string, error) {
	p := tea.NewProgram(initialActionModel(container))
	m, err := p.Run()
	if err != nil {
		return "", err
	}

	action := m.(actionChoice).selectedAction
	return action, nil
}