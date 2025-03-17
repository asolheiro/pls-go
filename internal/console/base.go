package console

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	palette "github.com/asolheiro/pls/internal/color-palette"
	grt "github.com/asolheiro/pls/internal/greetings"
	"github.com/asolheiro/pls/internal/settings"
	"github.com/asolheiro/pls/internal/tasks/operations"
	"github.com/asolheiro/pls/internal/utils"
	"github.com/charmbracelet/lipgloss"
	"github.com/jedib0t/go-pretty/text"

	"golang.org/x/term"
)

func PrintTasksTable(plt palette.ColorPalette) {
    sett, _ := settings.LoadConfigs()
    if sett.Quotes {
        grt.PrintGreeting(plt, sett.UserName)
        grt.PrintQuotes(plt)
    } else {
        grt.PrintGreeting(plt, sett.UserName)
    }

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

    colorPal := palette.NewPalette()
    
    header := fmt.Sprintf("%s%-*s%-*s%-*s",
        padStr,
        idWidth, colorPal.HeaderStyle.Render("ID   "),
        taskWidth, colorPal.HeaderStyle.Render("TASK                                                          "),
        statusWidth, colorPal.HeaderStyle.Render("STATUS"))
    
    divider := padStr + colorPal.DividerStyle.Render(strings.Repeat("━", tableWidth))
    
    fmt.Println(header)
    fmt.Println(divider)

    var (
        completedTasks int
        taskRow string
    )
    tasks, _ := operations.GetAllTasks()
    for index, task := range tasks {
        statusStyle := colorPal.TaskPendingStyle
        if task.Done {
            statusStyle = colorPal.TaskCompletedStyle
            completedTasks++
            
            indexStr := fmt.Sprintf("%d", index+1)
            
            indexRuneCount := utf8.RuneCountInString(indexStr)
            taskNameRuneCount := utf8.RuneCountInString(task.Name)
            
            indexStrStyled := colorPal.StrikeStyle.Render(indexStr)
            taskNameStyled := colorPal.StrikeStyle.Render(task.Name)
            
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
            statusChar = colorPal.StatusCharStyle.Render(MapDoneToChar(task.Done))
        } else {
            statusChar = statusStyle.Render(MapDoneToChar(task.Done))
        }

        taskRendered := statusStyle.Render(taskRow) + statusChar
        fmt.Println(taskRendered)
    }
        
    RenderProgressBar(plt, len(tasks), completedTasks)
}

func MapDoneToChar(done bool) string {
	if done {
		return "✓"
	} else {
		return "○"
	}
}

// RenderProgressBar creates a text-based progress bar.
func RenderProgressBar(plt palette.ColorPalette, total,  completed int) {
    remaingTasksBarStyle := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#3A3A3A")).
        Bold(true)
    taskCompletedStyle := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#F92672")).
        Bold(true)
    tasksRatioStyles := lipgloss.NewStyle().
        Foreground(lipgloss.Color("#86962E"))

    completedBar := taskCompletedStyle.Render(
        strings.Repeat("━", completed*5),
    )
    remaingTasksBar := remaingTasksBarStyle.Render(
        strings.Repeat("━", 5* (total - completed)),
    )
    taskRatio := tasksRatioStyles.Render(
        fmt.Sprintf("%d/%d", completed, total),
    )

    progressBar := plt.CompleteBar.Sprint(completedBar) + plt.FinishedBar.Sprintf(remaingTasksBar)
    completeBar := fmt.Sprintf("%s %s", progressBar, taskRatio)
    width, err := utils.GetTerminalFullWidth()
    if err != nil {
        fmt.Printf("error getting terminal width, err: %s", err)
    }
    centeredBar := text.Align(text.AlignCenter).Apply(completeBar, width)

    fmt.Println(centeredBar)
}