package greetings

import (
	"fmt"
	"log"
	"time"

	palette "github.com/asolheiro/pls/internal/color-palette"
	"github.com/asolheiro/pls/internal/settings"
	"github.com/asolheiro/pls/internal/utils"
	"github.com/charmbracelet/lipgloss"
	"github.com/jedib0t/go-pretty/text"
)

func PrintGreeting(plt palette.ColorStyles, name string) {
	rawGreeting := fmt.Sprintf("Hello, %s! %s", name, formatTimeAndDate(time.Now()))
	greetingFormatted := plt.HeaderGreetingStyle.
		Align(lipgloss.Center).
		Render(
			fmt.Sprintf("Hello, %s! %s", name, formatTimeAndDate(time.Now())),
		)

	linesPadd := plt.HeaderGreetingStyle.Render(utils.LinePaddings(rawGreeting))
	fmt.Println(
		linesPadd + " " + greetingFormatted + " " + linesPadd,
	)
}

func PrintQuotes(plt palette.ColorStyles) {
	size, _ := utils.GetTerminalFullWidth()
	quote, err := settings.GetRandQuote()
	if err != nil {
		log.Fatal(err)
	}

	quoteForm := text.AlignCenter.Apply(
		fmt.Sprintf("\"%s\"\n", quote.Content),
		size,
	)

	authorForm := text.AlignCenter.Apply(
		fmt.Sprintf(" ・ %s ・", quote.Author),
		size,
	)

	fmt.Println(
		plt.QuoteStyle.Render(quoteForm),
	)
	fmt.Println(
		plt.QuoteStyle.Render(authorForm),
	)
}

func formatTimeAndDate(t time.Time) string {
	return fmt.Sprintf("It's %d %s | %s",
		t.Day(),
		t.Format("Jan"),
		t.Format("03:04 PM"))
}
