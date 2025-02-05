package ui

import "fmt"

func RenderCursor(inLine bool) string {
	if inLine {
		render := fmt.Sprintf("\033[%sm>\033[0m", "32")
		return render
	}
	return " "
}

func RenderContainer(container string) string {
    return fmt.Sprintf("\033[%sm%s\033[0m", "38;2;214;112;214", container)
}

func RenderContainerSelected(container string, isSelected bool) string {
    if isSelected {
        return fmt.Sprintf("\033[%sm%s\033[0m", "32", container)
    }
    return container
}

func RenderActionSelected(action string, isSelected bool) string {
    if isSelected {
        return fmt.Sprintf("\033[%sm%s\033[0m", "32", action)
    }
    return action
}