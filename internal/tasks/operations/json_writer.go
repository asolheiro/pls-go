package operations

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func getFilePath(fileName string) string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".config", "pls", fileName)
}

// Helper function to write entries to JSON file
func writeJSONWithEntry(filename string, plsSet Settings) error {
	jsonData, err := json.MarshalIndent(plsSet, "", "  ")
	if err != nil {
		return fmt.Errorf("error creating JSON: %v", err)
	}
	
	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}
	
	return nil
}