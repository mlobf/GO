// First Line must be a Package Declation.
// To run a new Go program => " go run name_of_file.go"
// To create a new GO module ....
/*
	go mod init theNameOfTheApp
*/

package main

import "fmt"

func main() {
	//var whatToSay string
	//whatToSay = "Hello World, again"

	// Or, using a smarter syntax below
	whatToSay := "Hello World, again!!"

	sayHello(whatToSay)
}

func sayHello(whatToSay string) {
	fmt.Println(whatToSay)
}
