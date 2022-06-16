package main

/*
Now lets Go Learn function in
Go.
*/

import (
	"fmt"
)

// Time to build some structs

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
	fmt.Println("")
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("This buddy %s has a Total of %d legs.", a.Name, a.NumberOfLegs)
	fmt.Println("")
}

func main() {

	//f := addTwoNumbers(2, 2)
	//fmt.Println(f)
	fmt.Println("Lets Go learn how to create some functions in Go")

	s := sumMany(1, 2, 2, 3, 4, 4, 5, 90)
	fmt.Println(s)

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "au au"
	dog.NumberOfLegs = 4
	dog.Says()
	dog.HowManyLegs()

	var cat Animal
	cat.Name = "Tiao"
	cat.Sound = "Miau"
	dog.NumberOfLegs = 4
	cat.Says()
	cat.HowManyLegs()

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
