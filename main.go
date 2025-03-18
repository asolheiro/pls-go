package main

import (
	"fmt"
	"os"

	palette "github.com/asolheiro/pls/internal/color-palette"
	"github.com/asolheiro/pls/internal/console"
	"github.com/asolheiro/pls/internal/settings"
	"github.com/asolheiro/pls/internal/tasks/operations"
)

const fileName = "config.json"

func main() {
	settings.CheckAndCreateConfig(fileName)
	plt := palette.NewPalette()

	if len(os.Args) == 1 {
		console.RootCmd(plt)
	} else {
		switch os.Args[1] {
		case "add":
			if os.Args[2] == "" {
				fmt.Println("Insert a task")
			} else if os.Args[2] == "--help" {
				console.AddTaskHelp(plt)
			} else {
				if err := operations.AddTask(os.Args[2]); err == nil {
					console.AddTaskConsole(plt, os.Args[2])
				}
			}
		case "clean":
		case "clear":
		case "del":
		case "done":
		case "edit":
		case "move":
		case "tasks":
		case "callme":
		case "config":
		case "docs":
		case "quotes":
		case "setup":
		case "tasks-progress":
		case "version":
		case "cont-done":
		case "count-undone":
		case "--help":
			console.RootHelp(plt)
		default:
			console.CmdNotFound(plt)
		}

	}
}
