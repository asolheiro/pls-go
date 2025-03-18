package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"

	"golang.org/x/term"
)

func GetTerminalFullWidth() (int, error) {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	return width, err
}

func GetFilePath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".config", "pls", "config.json")
}

func ReplaceSpacesWithLines(s string) string {
	trimmed := strings.TrimSpace(s)

	totalWidth, _ := GetTerminalFullWidth()
	contentWidth := GetVisibleLength(trimmed)
	paddingWidth := (totalWidth - contentWidth - 2) / 2
	linePadding := strings.Repeat("─", paddingWidth)

	return fmt.Sprintf("%s %s %s", linePadding, trimmed, linePadding)
}

// Helper function to get the visible length of a string (excluding ANSI codes)
func GetVisibleLength(s string) int {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)

	cleanStr := re.ReplaceAllString(s, "")

	return utf8.RuneCountInString(cleanStr)
}

func LinePaddings(s string) string {
	trimmed := strings.TrimSpace(s)

	totalWidth, _ := GetTerminalFullWidth()
	contentWidth := GetVisibleLength(trimmed)
	paddingWidth := (totalWidth - contentWidth - 2) / 2
	linePadding := strings.Repeat("─", paddingWidth)
	return linePadding
}

func MapDoneToChar(done bool) string {
	if done {
		return "✓"
	} else {
		return "○"
	}
}
