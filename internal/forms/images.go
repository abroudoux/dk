package forms

import (
	"fmt"
	"strings"
	"time"

	"github.com/abroudoux/dk/internal/ui"
	"github.com/abroudoux/dk/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
	imagetypes "github.com/docker/docker/api/types/image"
)

type Image = imagetypes.Summary

type imageChoice struct {
	images []Image
	cursor int
	imageSelected Image
}

func initialImageModel(images []Image) imageChoice {
	return imageChoice{
		images: images,
		cursor: len(images) - 1,
		imageSelected: Image{},
	}
}

func (menu imageChoice) Init() tea.Cmd {
	return nil
}

func (menu imageChoice) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return menu, tea.Quit
		case "up":
			menu.cursor--
			if menu.cursor < 0 {
				menu.cursor = len(menu.images) - 1
			}
		case "down":
			menu.cursor++
			if menu.cursor >= len(menu.images) {
				menu.cursor = 0
			}
		case "enter":
			menu.imageSelected = menu.images[menu.cursor]
			return menu, tea.Quit
		}
	}

	return menu, nil
}

func (menu imageChoice) View() string {
    s := "\033[H\033[2J"
    s += "Choose an image:\n\n"

    for i, image := range menu.images {
        var imageLine string
        var name, tag string

        created := time.Unix(image.Created, 0).Format("2006-01-02 15:04:05")
        if len(image.RepoTags) > 0 {
            parts := strings.Split(image.RepoTags[0], ":")
            name = parts[0]
            if len(parts) > 1 {
                tag = parts[1]
            } else {
                tag = "latest"
            }
        } else {
            name = "<none>"
            tag = "<none>"
        }

        formattedSize := utils.FormatSize(image.Size)
        imageLine = fmt.Sprintf("%-30s %-10s %-12s %-20s %s", name, tag, image.ID[:20], created, formattedSize)

        cursor := " "
        cursor = ui.RenderCursor(menu.cursor == i)
        s += fmt.Sprintf("%s %s\n", cursor, ui.RenderSelected(imageLine, menu.cursor == i))
    }

    return s
}

func ChooseImage(images []Image) (Image, error) {
	p := tea.NewProgram(initialImageModel(images))
	m, err := p.Run()
	if err != nil {
		return Image{}, err
	}
	imageSelected := m.(imageChoice).imageSelected
	return imageSelected, nil
}
