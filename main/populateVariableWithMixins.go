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

func populateVariableWithMixins(token Token) Token {
tokenLoop:
	for {
		userChoice := getUserMixinChoice(token)
		var newValue string

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
				break
			}

			fmt.Println("Invalid input. Please try again")
			continue
		}

		token.values = append(token.values, newValue)
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
