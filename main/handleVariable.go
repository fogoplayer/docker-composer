package main

import (
	"fmt"
)

type UserChoice string

const (
	CREATE_NEW = "1"
	REUSE      = "2"
	CONTINUE   = "3"
	INVALID    = "INVALID"
)

// TODO allow template authors to provide defaults for tokens

func handleVariable(token Token) Token {
tokenLoop:
	for {
		userChoice := getUserChoice(token)

		switch userChoice {
		case CREATE_NEW:
			token.values = append(token.values, createMixin())

		case "2":
			fmt.Println(2)

		case "3":
			break tokenLoop

		default:
			fmt.Println("Invalid input. Please try again")
		}
	}
	return token
}

func getUserChoice(token Token) UserChoice {
	var userChoice UserChoice

	// Create prompt
	tokenName := token.name
	var defaultSelection UserChoice
	continueOption := ""
	allowContinue := len(token.values) > 0

	if allowContinue {
		defaultSelection = CONTINUE
		continueOption = `
	3) continue`
	} else {
		defaultSelection = REUSE
	}

	fmt.Printf(`Choose a behavior for %s (%s):
	1) create a new mixin
	2) reuse an existing mixin%s
`, tokenName, defaultSelection, continueOption)

	// process response
	bytes, _ := fmt.Scanf("%s", &userChoice)

	if bytes < 1 {
		userChoice = defaultSelection
	}

	if userChoice == "3" && !allowContinue {
		userChoice = INVALID
	}

	return userChoice
}
