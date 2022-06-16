package main

/*
Now lets Go Learn function in
Go.
*/

import (
	"fmt"
)

func main() {

	//f := addTwoNumbers(2, 2)
	//fmt.Println(f)
	fmt.Println("Lets Go learn how to create some functions in Go")

	s := sumMany(1, 2, 2, 3, 4, 4, 5, 90)
	fmt.Println(s)
}

// Sum two numbers
func addTwoNumbers(numOne, numTwo int) int {
	return numOne + numTwo
}

// Sum a array of numbers
// variatic params
func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}

	return total

}
