package main

import (
	"fmt"

	palette "github.com/asolheiro/pls/internal/color-palette"
	grt "github.com/asolheiro/pls/internal/greetings"
	"github.com/asolheiro/pls/internal/tasks/operations"
)

func main() {
	plt := palette.LoadColorPalette()

	grt.PrintGreeting(plt,"Armando")		

	if err := operations.ClearAllTask("config.json"); err != nil {
		fmt.Println(err)
	}
	
}