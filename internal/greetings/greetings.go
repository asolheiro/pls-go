package greetings

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	palette "github.com/asolheiro/pls/internal/color-palette"
	"github.com/asolheiro/pls/internal/settings"
	"github.com/jedib0t/go-pretty/text"
	"golang.org/x/term"
)

func PrintGreeting(plt palette.ColorPalette, name string) {
	size, _ := getTerminalFullWidth()
	greetingForm := text.AlignCenter.Apply(
		fmt.Sprintf("Hello, %s! %s", name, formatTimeAndDate(time.Now())), 
		size,
	)

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
	

	plt.HeaderGreeting.Println(replaceSpacesWithLines(greetingForm))
	plt.Quote.Println(quoteForm)
	plt.Author.Println(authorForm)
}

func getTerminalFullWidth() (int, error) {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	return width, err
}

func formatTimeAndDate(t time.Time) string {
	return fmt.Sprintf("It's %d %s | %s", 
		t.Day(),
		t.Format("Jan"),
		t.Format("03:04 PM"))
}

func replaceSpacesWithLines(s string) string {
	trimmed := strings.TrimSpace(s)
	
	totalWidth, _ := getTerminalFullWidth()
	contentWidth := len(trimmed)
	paddingWidth := (totalWidth - contentWidth - 2) / 2 // -2 for the spaces around the content
	
	linePadding := strings.Repeat("─", paddingWidth-1)
	
	return fmt.Sprintf("%s  %s  %s", linePadding, trimmed, linePadding)
}
