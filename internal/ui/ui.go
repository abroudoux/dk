package ui

import "fmt"

func RenderCursor(inLine bool) string {
	if inLine {
		render := fmt.Sprintf("\033[%sm>\033[0m", "32")
		return render
	}
	return " "
}

func RenderElementSelected(element string) string {
	return fmt.Sprintf("\033[%sm%s\033[0m", "38;2;214;112;214", element)
}

func RenderLineSelected(line string, isSelected bool) string {
    if isSelected {
        return fmt.Sprintf("\033[%sm%s\033[0m", "32", line)
    }
    return line
}