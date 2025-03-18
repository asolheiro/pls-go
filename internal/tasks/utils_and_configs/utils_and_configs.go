package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/asolheiro/pls/internal/tasks/operations"
	"github.com/asolheiro/pls/internal/utils"
)

func ChangeName(fileName string, newUserName string) error {
	filePath := utils.GetFilePath()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	var plsSettings operations.Settings
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	plsSettings.UserName = newUserName

	return operations.WriteJSONWithEntry(filePath, plsSettings)
}

func QuotesSwitch(fileName string) error {
	filePath := utils.GetFilePath()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	var plsSettings operations.Settings
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	plsSettings.Quotes = !plsSettings.Quotes

	return operations.WriteJSONWithEntry(filePath, plsSettings)
}

func ProgressSwitch(fileName string) error {
	filePath := utils.GetFilePath()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	var plsSettings operations.Settings
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	plsSettings.ShowTaskProgress = !plsSettings.ShowTaskProgress

	return operations.WriteJSONWithEntry(filePath, plsSettings)
}
