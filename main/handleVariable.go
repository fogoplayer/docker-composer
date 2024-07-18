package main

import (
	"fmt"
)

// TODO allow template authors to provide defaults for tokens

func handleVariable(token Token) Token {
	var userChoice string

while:
	for {
		fmt.Printf(`Choose a behavior for %s (2):
	1) create a new mixin
	2) reuse an existing mixin
`, token.name)
		bytes, _ := fmt.Scanf("%s", &userChoice)
		if bytes < 1 {
			userChoice = ""
		}

		switch userChoice {
		case "1":
			createMixin()
			break while

		case "":
			fallthrough

		case "2":
			fmt.Println(2)
			break while

		default:
			fmt.Println("Invalid input. Please try again")
		}
	}
	return token
}
