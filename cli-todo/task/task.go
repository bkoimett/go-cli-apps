package task

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Description string `json:"description"`
	Done		bool	`json:"done"`
}

const dataFile = "todo.json"

func loadTasks() ([]Task, error) {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	
	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	return tasks, err
}