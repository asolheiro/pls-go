package console

import (
	"fmt"

	palette "github.com/asolheiro/pls/internal/color-palette"
	"github.com/asolheiro/pls/internal/tasks/operations"
	"github.com/asolheiro/pls/internal/utils"
	"github.com/charmbracelet/lipgloss"
)

func AddTaskConsole(plt palette.ColorPalette, taskName string) {
	AddTaskHeader(plt, taskName)
	tasks, _ := operations.GetAllTasks()

	completedTasks := RenderTasksTable(tasks)
	RenderProgressBar(plt, len(tasks), completedTasks)
}

func AddTaskHeader(plt palette.ColorPalette, taskName string) {
	var (
		textStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A5D9A3")).
			Align(lipgloss.Center)
		linesStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#B38EE9"))
	)

	rawText := fmt.Sprintf(
		"Added %s to the list",
		taskName,
	)
	textFormatted := fmt.Sprintf(
		"Added %s to the list",
		textStyle.Render(taskName),
	)
	lineFormatted := linesStyle.Render(utils.LinePaddings(rawText))

	headerFormatted := fmt.Sprintf(
		lineFormatted + " " + textFormatted + " " + lineFormatted,
	)

	fmt.Println(headerFormatted)
}
