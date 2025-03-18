package console

import (
	"fmt"
	"os"

	palette "github.com/asolheiro/pls/internal/color-palette"
)

func CmdNotFound(plt palette.ColorStyles) {
	fmt.Printf(
		"\n%s%s\n",
		plt.HeaderGreetingStyle.
			PaddingLeft(2).
			Render("Usage:"),
		plt.QuoteStyle.
			Bold(true).
			PaddingLeft(1).
			Render("pls [OPTIONS] COMMAND [ARGS]...\n"),
	)
	cmd := plt.ErrorStyle.Render(os.Args[1])
	cmdNotFound := fmt.Sprintf("No such command '%s'", cmd)
	fmt.Println(
		DisplayBox(plt, cmdNotFound, "Error", len(cmdNotFound)),
	)
}
