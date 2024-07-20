package main

import (
	"fmt"
	"strconv"
)

type UserChoice string

const (
	CREATE_NEW UserChoice = "create a new mixin"
	REUSE      UserChoice = "reuse an existing mixin"
	CONTINUE   UserChoice = "continue"
	SKIP       UserChoice = "skip"
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

		case REUSE:
			userMixinChoice := getUserMixinChoice()
			numberToMixin := getListOfMixins()

			// if a number
			num, err := strconv.Atoi(string(userMixinChoice))
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

		case CONTINUE:
			fallthrough
		case SKIP:
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

	var moveOnString UserChoice
	if len(token.values) > 0 {
		moveOnString = CONTINUE
	} else {
		moveOnString = SKIP
	}

	return getUserSelection("Options for populating {{"+tokenName+"}}", createOptionMap(
		CREATE_NEW,
		REUSE,
		moveOnString,
	), "2")
}

func getUserMixinChoice() UserChoice {
	defaultChoice := UserChoice("1")

	fmt.Println("Saved mixins:")
	printSelectionList(getListOfMixins())

	fmt.Printf("Which would you like to choose? (%d): ", defaultChoice)

	return readLineFromStdInAsString(defaultChoice)
}
