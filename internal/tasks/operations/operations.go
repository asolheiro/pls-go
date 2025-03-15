package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

type Settings struct {
	UserName string`json:"user_name"`
	InitialSetupDone bool `json:"initial_setup_done"`
	ShowTaskProgress bool `json:"show_task_progress"`
	Quotes bool `json:"show_quotes"`
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Name string `json:"name"`
	Done bool `json:"done"`
}



func AddTask(fileName string, newTask Task) error {
	filePath := getFilePath(fileName)
    
    data, err := os.ReadFile(filePath)
    if err != nil {
        if os.IsNotExist(err) {
            // Pass the full file path here, not just the filename
            return writeJSONWithEntry(filePath, Settings{})
        }
        return fmt.Errorf("error reading file: %v", err)
    }
    
    var plsSettings Settings
    if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
        return fmt.Errorf("error parsing JSON: %v", err)
    }
    
    plsSettings.Tasks = append(plsSettings.Tasks, newTask)
    
    return writeJSONWithEntry(filePath, plsSettings)
}

func DeleteTask(fileName string, taskIndex int) error {
    filePath := getFilePath(fileName)
    
    data, err := os.ReadFile(filePath)
    if err != nil {
        return fmt.Errorf("error reading file: %v", err)
    }
    
    var plsSettings Settings
    if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
        return fmt.Errorf("error parsing JSON: %v", err)
    }
    
    // Check if the index is valid
    if taskIndex < 0 || taskIndex >= len(plsSettings.Tasks) {
        return fmt.Errorf("invalid task index: %d", taskIndex)
    }
    plsSettings.Tasks = slices.Delete(plsSettings.Tasks, taskIndex, taskIndex + 1)
    
    return writeJSONWithEntry(filePath, plsSettings)
}

func DoneTaks(fileName string, taskIndex int) error {
    filePath := getFilePath(fileName)
    
    data, err := os.ReadFile(filePath)
    if err != nil {
        return fmt.Errorf("error reading file: %v", err)
    }
    
    var plsSettings Settings
    if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
        return fmt.Errorf("error parsing JSON: %v", err)
    }
    
    // Check if the index is valid
    if taskIndex < 0 || taskIndex >= len(plsSettings.Tasks) {
        return fmt.Errorf("invalid task index: %d", taskIndex)
    }
	
	plsSettings.Tasks[taskIndex].Done = true

    return writeJSONWithEntry(filePath, plsSettings)
}

func CleanDoneTasks(fileName string) error {
    filePath := getFilePath(fileName)
    
    data, err := os.ReadFile(filePath)
    if err != nil {
        return fmt.Errorf("error reading file: %v", err)
    }
    
    var plsSettings Settings
    if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
        return fmt.Errorf("error parsing JSON: %v", err)
    }

    for i := len(plsSettings.Tasks) - 1; i >= 0; i-- {
        task := plsSettings.Tasks[i]
        if task.Done {
            plsSettings.Tasks = slices.Delete(plsSettings.Tasks, i, i + 1)
        }
    }

    return writeJSONWithEntry(filePath, plsSettings)
}

func ClearAllTask(fileName string) error {
    filePath := getFilePath(fileName)
    
    data, err := os.ReadFile(filePath)
    if err != nil {
        return fmt.Errorf("error reading file: %v", err)
    }
    
    var plsSettings Settings
    if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
        return fmt.Errorf("error parsing JSON: %v", err)
    }
    
    plsSettings.Tasks = slices.Delete(plsSettings.Tasks, 0, len(plsSettings.Tasks))

    return writeJSONWithEntry(filePath, plsSettings)
}

