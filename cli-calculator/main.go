package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bkoimett/go-cli-apps/cli-calculator/calculator"
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
		calculator.AddTwoNumbers(num1, num2)
	case "sub":
		calculator.SubtractTwoNumbers(num1, num2)
	case "div":
		if num1 > num2 {
			calculator.DivideTwoNumbers(num1, num2)
		} else {
			fmt.Println("📃Invert numbers: second number is higher than second")
		}
	case "multi":
		calculator.MultiplyTwoNumbers(num1, num2)
	case "mod":
		calculator.ModulusOfNumber(num1, num2)
	default:
		fmt.Println("⚠️ Illegal operation, use: <add> <sub> <div> <multi> <mod>")
	}
}
