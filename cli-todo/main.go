package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bkoimett/go-cli-apps/cli-todo/task"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: todo <add|list|done> [task description|task number]")
		return
	}

	tasks, err := task.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks: ", err)
		return
	}

	switch args[0] {
	case "list": 
		task.List(tasks)
	case "delete":
		if len(args) < 2 {
			fmt.Println("Please provide a task descrition.")
			return
		}
		description := args[1]
		tasks = task.Add(tasks, description)
		task.Save(tasks)
		fmt.Println("Task added.")
	case "done":
		if len(args) < 2 {
			fmt.Println("Please provide a task number.")
			return
		}
		num, err := strconv.Atoi(args[1])
		if err != nil || num < 1 || num > len(tasks) {
			fmt.Println("invalid task number.")
			return
		}
		task.MarkDone(tasks, num-1)
		task.Save(tasks)
		fmt.Println("Task marked as done")	
	default:
		fmt.Println("Unknown command.", args[0])
	}

	// try to log errors as well with time andstatus code
}