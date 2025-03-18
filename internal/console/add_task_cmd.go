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

func AddTaskHelp(plt palette.ColorStyles) {
	width, _ := utils.GetTerminalFullWidth()
	fmt.Printf(
		"\n%s%s\n%s\n",
		plt.TaskCompletedStyle.
			PaddingLeft(2).
			Render("Usage:"),
		plt.QuoteStyle.
			Bold(true).
			PaddingLeft(1).
			Render("pls add [OPTIONS] TASK\n"),
		plt.TaskCompletedStyle.
			PaddingLeft(2).
			Render("Add a Task ✨ (Add task name inside quotes)\n"),
	)
	ast := plt.ErrorStyle.Render("*")
	task := plt.QuoteStyle.Render("\ttask\t\tTEXT")
	defaults := plt.InfoStyle.Render("\t[default: None]")
	argumentsInfo := ast + task + defaults

	label := plt.InfoStyle.Render("Arguments")
	fmt.Println(
		DisplayBox(plt, argumentsInfo, label, width),
	)
	
	label = plt.InfoStyle.Render("Options")
	optionsCmdInfo := plt.QuoteStyle.
		Bold(true).
		PaddingLeft(2).
		Render("--help")
	optionsDescInfo := plt.InfoStyle.
		Render("\t\tShow this message and exit.")

	optionsInfo := optionsCmdInfo + optionsDescInfo

	fmt.Println(
		DisplayBox(plt, optionsInfo, label, width),
	)
}

func AddMissingTask(plt palette.ColorStyles) {
	fmt.Printf(
		"\n%s%s\n%s\n",
		plt.HeaderGreetingStyle.
			PaddingLeft(2).
			Render("Usage:"),
		plt.QuoteStyle.
			Bold(true).
			PaddingLeft(1).
			Render("pls add [OPTIONS] TASK\n"),
		plt.TaskCompletedStyle.
			PaddingLeft(2).
			Render("Try 'pls add --help' for help.\n"),
	)
	width, _ := utils.GetTerminalFullWidth()
	message := "Missing argument 'TASK'."
	errorLabel := plt.ErrorStyle.Bold(true).Render("Error")
	fmt.Println(
		DisplayBox(plt, message, errorLabel, width),
	)
}
