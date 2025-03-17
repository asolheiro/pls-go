package greetings

import (
	"fmt"
	"log"
	"time"

	palette "github.com/asolheiro/pls/internal/color-palette"
	"github.com/asolheiro/pls/internal/settings"
	"github.com/asolheiro/pls/internal/utils"
	"github.com/jedib0t/go-pretty/text"
)

func PrintGreeting(plt palette.ColorPalette, name string) {
	size, _ := utils.GetTerminalFullWidth()
	greetingForm := text.AlignCenter.Apply(
		fmt.Sprintf("Hello, %s! %s", name, formatTimeAndDate(time.Now())), 
		size,
	)
	plt.HeaderGreeting.Println(utils.ReplaceSpacesWithLines(greetingForm))
}

func PrintQuotes(plt palette.ColorPalette) {
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
	plt.Quote.Println(quoteForm)
	plt.Author.Println(authorForm)
}

func formatTimeAndDate(t time.Time) string {
	return fmt.Sprintf("It's %d %s | %s", 
		t.Day(),
		t.Format("Jan"),
		t.Format("03:04 PM"))
}


