// First Line must be a Package Declation.
// To run a new Go program => " go run name_of_file.go"
// To create a new GO module ....
/*
	go mod init theNameOfTheApp
*/

// To build a exe file using GO .....
// 	go build -o eliza main.go

package main

// Testing the git

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response((userInput)))
		}
	}
}
