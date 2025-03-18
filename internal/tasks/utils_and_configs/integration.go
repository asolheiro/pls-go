package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/asolheiro/pls/internal/tasks/operations"
	"github.com/asolheiro/pls/internal/utils"
)

func CountDone(fileName string) (int, error) {
	filePath := utils.GetFilePath()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	var plsSettings operations.Settings
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
		return 0, fmt.Errorf("error parsing JSON: %v", err)
	}

	var doneCounter int
	for _, task := range plsSettings.Tasks {
		if task.Done {
			doneCounter++
		}
	}
	return doneCounter, nil
}

func CountUndone(fileName string) (int, error) {
	filePath := utils.GetFilePath()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	var plsSettings operations.Settings
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
		return 0, fmt.Errorf("error parsing JSON: %v", err)
	}

	var undoneCounter int
	for _, task := range plsSettings.Tasks {
		if !task.Done {
			undoneCounter++
		}
	}
	return undoneCounter, nil
}
