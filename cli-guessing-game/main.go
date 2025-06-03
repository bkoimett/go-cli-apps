package main

import (
	"fmt"
	"math/rand"
)

func main() {
	guess := 0
	fmt.Println("Guess a number between 1-100")
	fmt.Scan(&guess)

	randomNum := rand.Intn(100)

	fmt.Println("The random number is :")
	fmt.Println(randomNum)
	fmt.Println()

	if guess == randomNum {
		fmt.Println("Your guess was correct")
	} else {
		fmt.Println("Your guess was incorrect")
	}
}
