package main

import (
	"bufio"
	"fmt"
	"os"
)

/*

	// one way - the long way
	var firstNumber int
	firstNumber = 2

	// another way - a new way
	var secondNumber = 2

	// subtration way, let GO figure out the type that should be used.
	subtraction := 7
*/

const prompt = "and press ENTER when ready"
const other = ""

func main() {

	var firstNumber = 2
	var secondNumber = 5
	var subtraction = 7
	var answer int

	reader := bufio.NewReader(os.Stdin)

	// Display a welcome/instructions.
	// Take them through the game.
	// give them the answer

	fmt.Println("Guess the Number")
	fmt.Println("================")
	fmt.Println("")
	fmt.Println("Think of a number between 1 and 10", prompt)

	reader.ReadString('\n')

	fmt.Println("Multiply your number by ", firstNumber, prompt)

	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)

	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)

	// give them answer.
	answer = firstNumber*secondNumber - subtraction

	fmt.Println("The answer is ", answer)
}
