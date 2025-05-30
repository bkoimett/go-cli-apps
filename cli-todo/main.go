package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 && os.Args[1] != "list" {
		fmt.Println("Caution: kindly use <operand> <taskid>/<task>")
		fmt.Println("⚠️")
		return
	}
	// create file in here but try to cretae it with a command line prompt
	op := os.Args[1]
	task := os.Args[2]

	switch op {
	case "add": 
		// write to list
	case "delete":
		// delete from file
	case "list":
		// show list - no need for task params
	default:
		fmt.Println()
	}

	// try to log errors as well with time andstatus code
}