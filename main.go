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
	plt := palette.LoadColorPalette()

	if len(os.Args) == 1{
		console.PrintTasksTable(plt)
	} else {
		switch os.Args[1] {
		case "add":
			if os.Args[2] != "" {
				if err := operations.AddTask(os.Args[2]); err == nil {
					console.AddTaskConsole(plt, os.Args[2])
				}
			} else {
				fmt.Println("Insert a task")
			}
		default:
			fmt.Println("Invalid command")
		}

	}
}


