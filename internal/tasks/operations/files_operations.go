package operations

import (
	"encoding/json"
	"fmt"
	"os"
)

// Helper function to write entries to JSON file
func WriteJSONWithEntry(filename string, plsSet Settings) error {
	jsonData, err := json.MarshalIndent(plsSet, "", "  ")
	if err != nil {
		return fmt.Errorf("error creating JSON: %v", err)
	}

	if err := os.WriteFile(filename, jsonData, 0644); err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}
