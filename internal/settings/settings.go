package settings

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type CodeSettings struct {
	configName           string
	configPath           string
	fullSettingsPath     string
	minimalDefaultConfig map[string]any
}



func NewSettings() *CodeSettings {
	s := &CodeSettings{}

	s.configName = s.getConfigName()
	s.configPath = s.getConfigPath()
	s.fullSettingsPath = filepath.Join(s.configPath, s.configName)
	s.minimalDefaultConfig = map[string]any{
		"user_name": "",
		"tasks":     []any{},
	}
	s.createDirIfNotExists()
	return s
}

func (s *CodeSettings) getConfigName() string {
	return "config.json"
}

func (s *CodeSettings) getConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error finding home diredtory, err: %s", err)
	}
	return filepath.Join(homeDir, ".config", "pls")
}

func (s *CodeSettings) GetFullSettingsPath() string {
	return s.fullSettingsPath
}

func (s *CodeSettings) createDirIfNotExists() {
	if _, err := os.Stat(s.configPath); os.IsNotExist(err) {
		os.MkdirAll(s.configPath, 0755)
	}
}

func (s *CodeSettings) ExistsSettings() bool {
	_, err := os.Stat(s.fullSettingsPath)
	return !os.IsNotExist(err)
}

func (s *CodeSettings) GetSettings() map[string]any {
	if _, err := os.Stat(s.fullSettingsPath); !os.IsNotExist(err) {
		file, err := os.Open(s.fullSettingsPath)
		if err != nil {
			return s.minimalDefaultConfig
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		settings := make(map[string]any)
		err = decoder.Decode(&settings)
		if err != nil {
			return s.minimalDefaultConfig
		}
		return settings
	}
	return s.minimalDefaultConfig
}

func (s *CodeSettings) WriteSettings(data map[string]any) error {
	file, err := os.Create(s.fullSettingsPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func (s *CodeSettings) GetName() string {
	settings := s.GetSettings()
	if name, ok := settings["user_name"].(string); ok {
		return name
	}
	return ""
}

func (s *CodeSettings) GetTasks() []map[string]any {
	settings := s.GetSettings()
	tasks, ok := settings["tasks"].([]any)
	if !ok {
		return []map[string]any{}
	}

	result := make([]map[string]any, len(tasks))
	for i, task := range tasks {
		if taskMap, ok := task.(map[string]any); ok {
			result[i] = taskMap
		}
	}
	return result
}

// ShowTasksProgress returns whether to show task progress
func (s *CodeSettings) ShowTasksProgress() bool {
	settings := s.GetSettings()
	if showProgress, ok := settings["show_task_progress"].(bool); ok {
		return showProgress
	}
	return true
}

func (s *CodeSettings) ShowQuotes() bool {
	settings := s.GetSettings()
	if showQuotes, ok := settings["show_quotes"].(bool); ok {
		return showQuotes
	}
	return true
}

func (s *CodeSettings) AllTasksDone() bool {
	tasks := s.GetTasks()
	if len(tasks) == 0 {
		return false
	}

	for _, task := range tasks {
		done, ok := task["done"].(bool)
		if !ok || !done {
			return false
		}
	}
	return true
}

func (s *CodeSettings) GetAllTasksUndone() []map[string]interface{} {
	tasks := s.GetTasks()
	result := []map[string]interface{}{}

	for _, task := range tasks {
		done, ok := task["done"].(bool)
		if !ok || !done {
			result = append(result, task)
		}
	}
	return result
}

func (s *CodeSettings) CountTasksDone() int {
	tasks := s.GetTasks()
	if len(tasks) == 0 {
		return 0
	}

	count := 0
	for _, task := range tasks {
		done, ok := task["done"].(bool)
		if ok && done {
			count++
		}
	}
	return count
}

func (s *CodeSettings) CountTasksUndone() int {
	tasks := s.GetTasks()
	if len(tasks) == 0 {
		return 0
	}

	count := 0
	for _, task := range tasks {
		done, ok := task["done"].(bool)
		if !ok || !done {
			count++
		}
	}
	return count
}
