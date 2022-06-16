package main

/*
Now lets Go Learn function in
Go.
*/

import (
	"fmt"
)

func main() {

	f := addTwoNumbers(2, 2)
	fmt.Println(f)

	fmt.Println("Lets Go learn how to create some functions in Go")

}

func addTwoNumbers(numOne, numTwo int) int {
	return numOne + numTwo
}
