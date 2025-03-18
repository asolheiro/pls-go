package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"slices"

	"github.com/asolheiro/pls/internal/utils"
)

type Settings struct {
	UserName         string `json:"user_name"`
	InitialSetupDone bool   `json:"initial_setup_done"`
	ShowTaskProgress bool   `json:"show_task_progress"`
	Quotes           bool   `json:"show_quotes"`
	Tasks            []Task `json:"tasks"`
}

type Task struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func AddTask(newTaskName string) error {
	configPath := utils.GetFilePath()

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return WriteJSONWithEntry(configPath, Settings{})
		}
		return fmt.Errorf("error reading file: %v", err)
	}

	var plsSettings Settings
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	plsSettings.Tasks = append(
		plsSettings.Tasks,
		Task{
			newTaskName,
			false,
		})

	return WriteJSONWithEntry(configPath, plsSettings)
}

func DeleteTask(string, taskIndex int) error {
	configPath := utils.GetFilePath()

	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	var plsSettings Settings
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	if taskIndex < 0 || taskIndex >= len(plsSettings.Tasks) {
		return fmt.Errorf("invalid task index: %d", taskIndex)
	}
	plsSettings.Tasks = slices.Delete(plsSettings.Tasks, taskIndex, taskIndex+1)

	return WriteJSONWithEntry(configPath, plsSettings)
}

func MarkAsDoneTaks(string, taskIndex int) error {
	configPath := utils.GetFilePath()

	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	var plsSettings Settings
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	if taskIndex < 0 || taskIndex >= len(plsSettings.Tasks) {
		return fmt.Errorf("invalid task index: %d", taskIndex)
	}

	plsSettings.Tasks[taskIndex].Done = true

	return WriteJSONWithEntry(configPath, plsSettings)
}

func MarkAsUndoneTaks(string, taskIndex int) error {
	configPath := utils.GetFilePath()

	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	var plsSettings Settings
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	if taskIndex < 0 || taskIndex >= len(plsSettings.Tasks) {
		return fmt.Errorf("invalid task index: %d", taskIndex)
	}

	plsSettings.Tasks[taskIndex].Done = false

	return WriteJSONWithEntry(configPath, plsSettings)
}

func EditTaskName(string, taskIndex int, newName string) error {
	configPath := utils.GetFilePath()

	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	var plsSettings Settings
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	if taskIndex < 0 || taskIndex >= len(plsSettings.Tasks) {
		return fmt.Errorf("invalid task index: %d", taskIndex)
	}

	plsSettings.Tasks[taskIndex].Name = newName

	return WriteJSONWithEntry(configPath, plsSettings)
}

func CleanDoneTasks(string) error {
	configPath := utils.GetFilePath()

	data, err := os.ReadFile(configPath)
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
			plsSettings.Tasks = slices.Delete(plsSettings.Tasks, i, i+1)
		}
	}

	return WriteJSONWithEntry(configPath, plsSettings)
}

func ClearAllTask(string) error {
	configPath := utils.GetFilePath()

	data, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	var plsSettings Settings
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	plsSettings.Tasks = slices.Delete(plsSettings.Tasks, 0, len(plsSettings.Tasks))

	return WriteJSONWithEntry(configPath, plsSettings)
}

func GetAllTasks() ([]Task, error) {
	configPath := utils.GetFilePath()

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var plsSettings Settings
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&plsSettings); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return plsSettings.Tasks, nil
}
