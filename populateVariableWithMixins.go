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

func populateVariableWithMixins(variableName string, providedValues ...UserChoice) []string {
	var newValue string
	userChoice := INVALID
	cliMode := len(providedValues) > 0
	contents := []string{}

	if len(providedValues) > 0 {
		userChoice = providedValues[0]
		println("Populating {{"+variableName+"}} with mixin", userChoice)
	}

tokenLoop:
	for {
		if userChoice == INVALID {
			userChoice = getUserMixinChoice(variableName, len(contents) > 0)
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

		contents = append(contents, newValue)
		userChoice = INVALID

		if cliMode {
			break tokenLoop
		}
	}

	return contents
}

func getUserMixinChoice(tokenName string, hasBeenPopulated bool) UserChoice {
	// Create prompt
	var moveOnString UserChoice
	var defaultChoice UserChoice
	if hasBeenPopulated {
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
