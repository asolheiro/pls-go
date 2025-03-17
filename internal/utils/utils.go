package utils

import (
	"os"
	"path/filepath"

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