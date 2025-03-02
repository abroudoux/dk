package images

import (
	"fmt"

	"github.com/abroudoux/dk/internal/ui"
	"github.com/abroudoux/dk/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	imageActionExit imageAction = iota
	imageActionDelete
	imageActionRun
)

func (a imageAction) String() string {
	return [...]string{"Exit", "Delete", "Run"}[a]
}

func initialActionModel(image image) actionChoice {
	actions := []imageAction{
		imageActionExit,
		imageActionDelete,
		imageActionRun,
	}

	return actionChoice{
		actions:        actions,
		cursor:         len(actions) - 1,
		actionSelected: imageActionExit,
		imageSelected:  image,
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
			menu.actionSelected = imageActionExit
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
	s += fmt.Sprintf("Image: %s\n\n", utils.RenderImageName(menu.imageSelected))

	for i, action := range menu.actions {
		cursor := " "
		cursor = ui.RenderCursor(menu.cursor == i)
		s += fmt.Sprintf("%s %s\n", cursor, ui.RenderLineSelected(action.String(), menu.cursor == i))
	}

	s += "\n"

	return s
}

func selectAction(image image) (imageAction, error) {
	p := tea.NewProgram(initialActionModel(image))
	m, err := p.Run()
	if err != nil {
		return imageActionExit, err
	}

	action := m.(actionChoice).actionSelected
	return action, nil
}
