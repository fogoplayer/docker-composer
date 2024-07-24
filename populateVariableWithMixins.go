package main

import (
	"fmt"
	"strconv"
)

type UserChoice string

const (
	CREATE_NEW UserChoice = "create a new mixin"
	CONTINUE   UserChoice = "continue"
	SKIP       UserChoice = "skip"
	INVALID    UserChoice = "INVALID"
)

// TODO allow template authors to provide defaults for tokens

func populateVariableWithMixins(token Token, providedValues ...UserChoice) Token {
	var newValue string
	userChoice := INVALID
	cliMode := len(providedValues) > 0

	if len(providedValues) > 0 {
		userChoice = providedValues[0]
		println("Populating {{"+token.name+"}} with mixin", userChoice)
	}

tokenLoop:
	for {
		if userChoice == INVALID {
			userChoice = getUserMixinChoice(token)
		}

		switch userChoice {
		case CREATE_NEW:
			var e error
			newValue, e = createMixin()
			if e != nil {
				panic(e)
			}

		case CONTINUE:
			fallthrough
		case SKIP:
			break tokenLoop

		default:
			mixin, e := getMixinContents(userChoice)
			if e == nil {
				newValue = mixin
				break // switch
			}

			fmt.Println("Invalid input. Please try again")
			userChoice = INVALID
			continue
		}

		token.values = append(token.values, newValue)
		userChoice = INVALID

		if cliMode {
			break tokenLoop
		}
	}

	return token
}

func getUserMixinChoice(token Token) UserChoice {
	// Create prompt
	tokenName := token.name

	var moveOnString UserChoice
	var defaultChoice UserChoice
	if len(token.values) > 0 {
		moveOnString = CONTINUE
		defaultChoice = UserChoice(strconv.Itoa(len(getListOfMixins()) + 2)) // because we append two options
	} else {
		moveOnString = SKIP
		defaultChoice = "1"
	}

	return getUserSelection(
		"Options for populating {{"+tokenName+"}}",
		append(
			getListOfMixins(),
			CREATE_NEW,
			moveOnString,
		),
		defaultChoice,
	)
}
