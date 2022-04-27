package main

import (
	"fmt"
)

func main() {

	// one way - the long way
	var firstNumber int
	firstNumber = 2

	// another way - a new way
	var secondNumber = 2

	// subtration way, let GO decides witch type should be used at that variable.
	subtraction := 7

	fmt.Print(firstNumber, secondNumber, subtraction)
}
