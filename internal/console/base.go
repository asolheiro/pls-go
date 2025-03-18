package console

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	palette "github.com/asolheiro/pls/internal/color-palette"
	"github.com/asolheiro/pls/internal/tasks/operations"
	"github.com/asolheiro/pls/internal/utils"
	"github.com/charmbracelet/lipgloss"
	"github.com/jedib0t/go-pretty/text"

	"golang.org/x/term"
)

// RenderTasksTable creates a text-based table to show tasks
func RenderTasksTable(plt palette.ColorStyles, tasks []operations.Task) int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80
	}

	idWidth := 4
	statusWidth := 8
	taskWidth := 80 - idWidth - statusWidth - 4
	tableWidth := 80
	leftPadding := (width - tableWidth) / 2
	if leftPadding < 0 {
		leftPadding = 0
	}
	padStr := strings.Repeat(" ", leftPadding)

	header := fmt.Sprintf("%s%-*s%-*s%-*s",
		padStr,
		idWidth, plt.HeaderStyle.Render("ID   "),
		taskWidth, plt.HeaderStyle.Render("TASK                                                          "),
		statusWidth, plt.HeaderStyle.Render("STATUS"))

	divider := padStr + plt.DividerStyle.Render(strings.Repeat("━", tableWidth))

	fmt.Println(header)
	fmt.Println(divider)

	var (
		completedTasks int
		taskRow        string
	)

	for index, task := range tasks {
		statusStyle := plt.TaskPendingStyle
		if task.Done {
			statusStyle = plt.TaskCompletedStyle
			completedTasks++

			indexStr := fmt.Sprintf("%d", index+1)

			indexRuneCount := utf8.RuneCountInString(indexStr)
			taskNameRuneCount := utf8.RuneCountInString(task.Name)

			indexStrStyled := plt.StrikeStyle.Render(indexStr)
			taskNameStyled := plt.StrikeStyle.Render(task.Name)

			taskRow = fmt.Sprintf("%s %s%s %s%s",
				padStr,
				indexStrStyled,
				strings.Repeat(" ", idWidth-1-indexRuneCount),
				taskNameStyled,
				strings.Repeat(" ", taskWidth-1-taskNameRuneCount),
			)
		} else {
			indexStr := fmt.Sprintf("%d", index+1)
			indexRuneCount := utf8.RuneCountInString(indexStr)
			taskNameRuneCount := utf8.RuneCountInString(task.Name)

			taskRow = fmt.Sprintf("%s %s%s %s%s",
				padStr,
				indexStr,
				strings.Repeat(" ", idWidth-1-indexRuneCount),
				task.Name,
				strings.Repeat(" ", taskWidth-1-taskNameRuneCount),
			)
		}
		var statusChar string
		if task.Done {
			statusChar = plt.StatusCharStyle.Render(utils.MapDoneToChar(task.Done))
		} else {
			statusChar = statusStyle.Render(utils.MapDoneToChar(task.Done))
		}

		taskRendered := statusStyle.Render(taskRow) + statusChar
		fmt.Println(taskRendered)
	}
	return completedTasks
}

// RenderProgressBar creates a text-based progress bar.
func RenderProgressBar(plt palette.ColorStyles, total, completed int) {
	completedTasksBar := plt.CompletedsBar.Render(
		strings.Repeat("━", 3*completed),
	)
	remaingTasksBar := plt.RemaingTasksBarStyle.Render(
		strings.Repeat("━", 3*(total-completed)),
	)
	taskRatio := plt.TaskRatioStyles.Render(
		fmt.Sprintf("%d/%d", completed, total),
	)

	width, err := utils.GetTerminalFullWidth()
	if err != nil {
		fmt.Printf("error getting terminal width, err: %s", err)
	}

	progressBar := fmt.Sprintf("%s%s %s", completedTasksBar, remaingTasksBar, taskRatio)
	fmt.Println(
		text.
			Align(text.AlignCenter).
			Apply(progressBar, width),
	)
}

// DisplayBox creates a box with a label and message, handling styled text properly
func DisplayBox(plt palette.ColorStyles, message, label string, width int) string {
	var boxStyle lipgloss.Style
	if label == "Error" {
		boxStyle = plt.ErrorStyle
	}

	labelFormatted := boxStyle.Render(label)

	topLeft := "╭"
	topRight := "╮"
	bottomLeft := "╰"
	bottomRight := "╯"
	horizontal := "─"
	vertical := "│"

	visibleLabelLen := lipgloss.Width(labelFormatted)
	visibleMsgLen := lipgloss.Width(message)

	contentWidth := width - 4

	labelWithSpaces := " " + labelFormatted + " "
	visibleLabelWithSpacesLen := visibleLabelLen + 2
	topBorder := topLeft + horizontal + labelWithSpaces + strings.Repeat(horizontal, width-visibleLabelWithSpacesLen-3) + topRight

	paddingLen := contentWidth - visibleMsgLen
	paddedMessage := " " + message + strings.Repeat(" ", paddingLen) + " "
	messageLine := vertical + paddedMessage + vertical

	bottomBorder := bottomLeft + strings.Repeat(horizontal, width-2) + bottomRight

	return topBorder + "\n" + messageLine + "\n" + bottomBorder
}
