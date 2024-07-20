package main

import (
	"fmt"
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
			userMixinChoice := getUserSelection(
				"Choose a saved mixin:",
				getListOfMixins(),
			)

			mixin, err := getMixinContents(userMixinChoice)
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
			continue
		}

		token.values = append(token.values, newValue)
	}
	return token
}

func getUserHandleVariableChoice(token Token) UserChoice {
	// Create prompt
	tokenName := token.name

	var moveOnString UserChoice
	var defaultChoice UserChoice
	if len(token.values) > 0 {
		moveOnString = CONTINUE
		defaultChoice = "3"
	} else {
		moveOnString = SKIP
		defaultChoice = "2"
	}

	return getUserSelection(
		"Options for populating {{"+tokenName+"}}",
		createOptionMap(
			CREATE_NEW,
			REUSE,
			moveOnString,
		),
		defaultChoice,
	)
}
