package main

import (
	"encoding/json"
	"os"
)

const tasksFile = "tasks.json"

// SaveTasks guarda una lista de tareas en un archivo JSON.
func SaveTask(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(tasksFile, data, 0644)
}

func LoadTasks(filename string) ([]Task, error) {
	data, err := os.ReadFile(tasksFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}
