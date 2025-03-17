package console

import (
	palette "github.com/asolheiro/pls/internal/color-palette"
	grt "github.com/asolheiro/pls/internal/greetings"
	"github.com/asolheiro/pls/internal/settings"
	op "github.com/asolheiro/pls/internal/tasks/operations"
)

func PrintTasksTable(plt palette.ColorPalette) {
    sett, _ := settings.LoadConfigs()
    if sett.Quotes {
        grt.PrintGreeting(plt, sett.UserName)
        grt.PrintQuotes(plt)
    } else {
        grt.PrintGreeting(plt, sett.UserName)
    }

    tasks, _ := op.GetAllTasks()
    completedTasks := RenderTasksTable(tasks)
        
    RenderProgressBar(plt, len(tasks), completedTasks)
}