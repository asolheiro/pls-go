package console

import (
	"fmt"

	palette "github.com/asolheiro/pls/internal/color-palette"
	grt "github.com/asolheiro/pls/internal/greetings"
	"github.com/asolheiro/pls/internal/settings"
	op "github.com/asolheiro/pls/internal/tasks/operations"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
	commandStyle = lipgloss.NewStyle().
			Foreground(palette.GraySilver).
			Bold(true)
	descriptionStyle = lipgloss.NewStyle().
				Foreground(palette.GraySilver).
				MarginRight(1)

	headerStyle = lipgloss.NewStyle().
			Foreground(palette.GrayAsh).
			Bold(true).
			Align(lipgloss.Center)
	cellStyle = lipgloss.NewStyle().Padding(0, 1).Width(55)
)

func RootCmd(plt palette.ColorStyles) {
	sett, _ := settings.LoadConfigs()
	if sett.Quotes {
		grt.PrintGreeting(plt, sett.UserName)
		grt.PrintQuotes(plt)
	} else {
		grt.PrintGreeting(plt, sett.UserName)
	}

	tasks, _ := op.GetAllTasks()
	completedTasks := RenderTasksTable(plt, tasks)

	RenderProgressBar(plt, len(tasks), completedTasks)
}

func RootHelp(plt palette.ColorStyles) {
	fmt.Printf(
		"\n%s%s\n",
		plt.TaskCompletedStyle.
			PaddingLeft(2).
			Render("Usage:"),
		plt.QuoteStyle.
			Bold(true).
			PaddingLeft(1).
			Render("pls [OPTIONS] COMMAND [ARGS]..."),
	)

	fmt.Printf(
		"\n%s\n%s\n",
		plt.QuoteStyle.
			PaddingLeft(2).
			Bold(true).
			Render("💻 PLS-CLI"),
		plt.TaskCompletedStyle.
			PaddingLeft(2).
			Render("・Minimalist and full configurable greetings and TODO list・"),
	)

	cmdsTb := commandsTable()
	utilsTb := utilsConfigsTable()
	optionsTb := optionsTable()

	fmt.Println(optionsTb)
	fmt.Println(cmdsTb)
	fmt.Println(utilsTb)
}

func optionsTable() *table.Table {
	optionsRows := [][]string{
		{commandStyle.Render("--install-completion"), descriptionStyle.Render("- Install completion for the current shell.")},
		{commandStyle.Render("--show-completion"), descriptionStyle.Render("- Show completion for the current shell, to copy it or customize the installation.")},
		{commandStyle.Render("--help"), descriptionStyle.Render("- Show this message and exit.")},
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().
			Foreground(palette.GrayAsh)).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return headerStyle
			}
			return cellStyle
		}).
		Headers("Utils & Configs", "Description").
		Rows(optionsRows...)
	return t
}

func commandsTable() *table.Table {
	commandsRows := [][]string{
		{commandStyle.Render("✨ add"), descriptionStyle.Render("- Add a Task (Add task name inside quotes)")},
		{commandStyle.Render("🧹 clean"), descriptionStyle.Render("- Clean up tasks marked as done")},
		{commandStyle.Render("🗑  clear"), descriptionStyle.Render("- Clear all tasks ")},
		{commandStyle.Render("✂️ del"), descriptionStyle.Render("- Delete a task")},
		{commandStyle.Render("👍🏽 done"), descriptionStyle.Render("- Mark a task as done ")},
		{commandStyle.Render("✏️ edit"), descriptionStyle.Render("- Edit a task by id (Add task name inside quotes)")},
		{commandStyle.Render("🔀 move"), descriptionStyle.Render("- Change task order")},
		{commandStyle.Render("📖 tasks"), descriptionStyle.Render("- Show all Tasks")},
		{commandStyle.Render("👎🏽 undone"), descriptionStyle.Render("- Mark a task as undone")},
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(palette.GrayAsh)).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return headerStyle
			}
			return cellStyle
		}).
		Headers("Options", "Description").
		Rows(commandsRows...)
	return t
}

func utilsConfigsTable() *table.Table {
	utilsRows := [][]string{
		{commandStyle.Render("🔊 callme"), descriptionStyle.Render("- Change name 📛 (without resetting data)")},
		{commandStyle.Render("📂 config"), descriptionStyle.Render("- Launch config directory")},
		{commandStyle.Render("🌐 docs"), descriptionStyle.Render("- Launch docs Website")},
		{commandStyle.Render("🏷  quotes"), descriptionStyle.Render("- Show quotes")},
		{commandStyle.Render("🔧 setup"), descriptionStyle.Render("- Reset all data and run setup")},
		{commandStyle.Render("🎯 tasks-progress"), descriptionStyle.Render("- Show tasks progress")},
		{commandStyle.Render("🔖 version"), descriptionStyle.Render("- Show version")},
	}

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().
			Foreground(palette.GrayAsh)).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return headerStyle
			}
			return cellStyle
		}).
		Headers("Utils & Configs", "Description").
		Rows(utilsRows...)
	return t
}
