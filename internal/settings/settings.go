package settings

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/asolheiro/pls/internal/tasks/operations"
	"github.com/asolheiro/pls/internal/utils"
)


func CheckAndCreateConfig(fileName string) error {
	configPath := utils.GetFilePath()
    _, err := os.Stat(configPath)
    if err == nil {
		return nil
    }
    
    if !os.IsNotExist(err) {
		return fmt.Errorf("error checking config file: %v", err)
    }

    homeDir, err := os.UserHomeDir()
    if err != nil {
        return fmt.Errorf("error getting home directory: %v", err)
    }
	configDir := filepath.Join(homeDir, ".config", "pls", )
    if err := os.MkdirAll(configDir, 0755); err != nil {
        return fmt.Errorf("error creating config directory: %v", err)
    }
    
    // Create default settings
    defaultSettings := operations.Settings{
        UserName:         "",
        InitialSetupDone: false,
        ShowTaskProgress: true,
        Quotes:           true,
        Tasks:            []operations.Task{},
    }
    
    jsonData, err := json.MarshalIndent(defaultSettings, "", "  ")
    if err != nil {
        return fmt.Errorf("error creating JSON: %v", err)
    }
    
    if err := os.WriteFile(configPath, jsonData, 0644); err != nil {
        return fmt.Errorf("error writing config file: %v", err)
    }
    
    return nil
}

func LoadConfigs() (operations.Settings, error) {
	configPath := utils.GetFilePath()
	
	data, err := os.ReadFile(configPath)
	if err != nil {
		return operations.Settings{}, fmt.Errorf("error reading file: %v", err)
	}
	
	var plsSettings operations.Settings
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
		return operations.Settings{}, fmt.Errorf("error parsing JSON: %v", err)
	}

	return plsSettings, nil
}

