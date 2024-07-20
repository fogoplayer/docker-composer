package main

import (
	"fmt"
	"strconv"
)

type UserChoice string

const (
	CREATE_NEW UserChoice = "1"
	REUSE      UserChoice = "2"
	CONTINUE   UserChoice = "3"
	INVALID    UserChoice = "INVALID"
)

// TODO allow template authors to provide defaults for tokens

func handleVariable(token Token) Token {
tokenLoop:
	for {
		userChoice := getUserHandleVariableChoice(token)
		var newValue string

		switch userChoice {
		case CREATE_NEW:
			newValue, _ = createMixin()

		case "2":
			userMixinChoice := getUserMixinChoice()
			numberToMixin := listMixins()

			// if a number
			num, err := strconv.Atoi(userMixinChoice)
			if err == nil {
				mixin, err := getMixin(numberToMixin[num])
				if err != nil {
					fmt.Println("Unable to retrieve mixin. Please try again.")
					break
				}
				newValue = mixin
				break
			}

			// if a full string
			mixin, err := getMixin(userMixinChoice)
			if err != nil {
				fmt.Println("Unable to retrieve mixin. Please try again.")
				break
			}
			newValue = mixin

		case "3":
			break tokenLoop

		default:
			fmt.Println("Invalid input. Please try again")
		}

		token.values = append(token.values, newValue)
	}
	return token
}

func getUserHandleVariableChoice(token Token) UserChoice {

	// Create prompt
	tokenName := token.name
	var defaultSelection UserChoice
	continueOption := ""
	allowContinue := len(token.values) > 0

	if allowContinue {
		defaultSelection = CONTINUE
		continueOption = "\n\t3) continue"
	} else {
		defaultSelection = REUSE
	}

	fmt.Printf(`Options for  {{%s}}:
	1) create a new mixin
	2) reuse an existing mixin%s
How would you like to populate {{%s}}? (%s): `, tokenName, continueOption, tokenName, defaultSelection)

	// process response
	userChoice := UserChoice(readLineFromStdInAsString(string(defaultSelection)))

	if userChoice == "3" && !allowContinue {
		userChoice = INVALID
	}

	return userChoice
}

func getUserMixinChoice() string {
	defaultChoice := 1

	fmt.Println("Saved mixins:")
	printSelectionList(listMixins())

	fmt.Printf("Which would you like to choose? (%d): ", defaultChoice)

	return readLineFromStdInAsString(string(defaultChoice))
}
