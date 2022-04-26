// First Line must be a Package Declation.
// To run a new Go program => " go run name_of_file.go"
// To create a new GO module ....
/*
	go mod init theNameOfTheApp
*/

package main

// Testing the git

import (
	"fmt"
	"myapp/doctor"
)

func main() {

	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

}
