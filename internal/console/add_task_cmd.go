package console

import (
	"fmt"

	palette "github.com/asolheiro/pls/internal/color-palette"
	"github.com/asolheiro/pls/internal/tasks/operations"
	"github.com/asolheiro/pls/internal/utils"
	"github.com/charmbracelet/lipgloss"
)

func AddTaskConsole(plt palette.ColorStyles, taskName string) {
	AddTaskHeader(plt, taskName)
	tasks, _ := operations.GetAllTasks()

	completedTasks := RenderTasksTable(plt, tasks)
	RenderProgressBar(plt, len(tasks), completedTasks)
}

func AddTaskHeader(plt palette.ColorStyles, taskName string) {
	var (
		textStyle = lipgloss.NewStyle().
				Foreground(palette.GreenMint).
				Align(lipgloss.Center)
		linesStyle = lipgloss.NewStyle().
				Foreground(palette.PurpleLavender)
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

func AddTaskHelp() {

}
