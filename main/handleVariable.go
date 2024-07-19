package main

import (
	"fmt"
	"strconv"
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
		userChoice := getUserMainMenuChoice(token)
		var newValue string

		switch userChoice {
		case CREATE_NEW:
			newValue = createMixin()

		case "2":
			newValue = getUserMixinChoice()

		case "3":
			break tokenLoop

		default:
			fmt.Println("Invalid input. Please try again")
		}

		token.values = append(token.values, newValue)
	}
	return token
}

func getUserMainMenuChoice(token Token) UserChoice {
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

func getUserMixinChoice() string {
	var userChoice string
	defaultChoice := 1

	fmt.Printf("Choose a mixin (%d):\n", defaultChoice)
	numberToMixin := listMixins()

	/* bytes, err :=  */
	fmt.Scanln(&userChoice)

	// if default
	if userChoice == "" {
		return getMixin(numberToMixin[defaultChoice])
	}

	// if a number
	num, err := strconv.Atoi(userChoice)
	if err == nil {
		return getMixin(numberToMixin[num])
	}

	// if a full string
	return getMixin(userChoice)
}
