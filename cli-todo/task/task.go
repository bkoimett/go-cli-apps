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

func LoadTasks() ([]Task, error) {
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

// write tasks back to the json file
func Save(task []Task) error {
	data, err := json.Marshal(task, "", "")
	if err != nil {
		return err
	}

	return os.WriteFile(dataFile, data, 0644) // owner, u
}

func List(task []Task) {
	if len(task) == 0 {
		fmt.Println("No tasks")
		return
	}

	for i, t := range task {
		status := "[ ]"
		if t.Done {
			status = "[X]"
		}
		fmt.Println("%d. %s %s\n", 1+i, status, t.Description)
	}
}