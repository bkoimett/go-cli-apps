package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("🤯Not enough parameters")
		fmt.Println("📃USAGE: go run . <operation> <num1> <num2>")
		return
	}

	op := os.Args[1]
	num1, _ := strconv.Atoi(os.Args[2])
	num2, _ := strconv.Atoi(os.Args[3])
	
	fmt.Println(op, num1, num2)

	switch op {
	case "add":
		addTwoNumbers(num1, num2)
	case "sub": 
		subtractTwoNumbers(num1, num2)
	case "div":
		if num1 > num2 {
			divideTwoNumbers(num1, num2)
		} else {
			fmt.Println("📃Invert numbers: second number is higher than second")
		}
	case "multi":
		multiplyTwoNumbers(num1, num2)
	case "mod":
		modulusOfNumber(num1, num2)
	default:
		fmt.Println("⚠️ Illegal operation, use: <add> <sub> <div> <multi> <mod>")
	}
}