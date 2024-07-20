package main

import (
	"fmt"
	"os"
)

const (
	BUILD_DOCKERFILE UserChoice = "build dockerfile from template"
	MANAGE_TEMPLATES UserChoice = "manage templates"
	MANAGE_MIXINS    UserChoice = "manage mixins"
	EXIT             UserChoice = "exit"
)

var home string = os.Getenv("HOME")
var contentPath Path = segmentsToPath(home, ".config", "docker-composer")

// Entry point for docker-composer
func main() {
	for {
		userChoice := getUserMainMenuChoice()

		switch userChoice {
		case BUILD_DOCKERFILE:
			buildDockerfileMenuOption()

		case MANAGE_TEMPLATES:
			manageTemplatesMenuOption()

		case MANAGE_MIXINS:
			manageMixinsMenuOption()

		case EXIT:
			os.Exit(0)

		default:
			fmt.Println("Invalid input. Please try again")
			continue
		}
	}
}

// Initial prompt when the user opens the application
func getUserMainMenuChoice() UserChoice {
	userChoice := getUserSelection(
		"What would you like to do?:",
		createOptionMap(
			BUILD_DOCKERFILE,
			MANAGE_TEMPLATES,
			MANAGE_MIXINS,
			EXIT,
		),
	)

	return userChoice
}

// Code to execute if the user chooses to build a dockerfile.
//
// Reads in a template, tokenizes it, replaces variables, and saves to a directory of the user's choice
func buildDockerfileMenuOption() {
	template, _ := readStringFromFile("/home/zarinloosli/docker-composer/example.tplt")
	ast := tokenize(template)
	for i, token := range ast {
		if token.kind == VARIABLE {
			ast[i] = handleVariable(token)
		}
	}

	dockerfile := buildDockerfileFromAst(ast)
	saveDockerFile(dockerfile)
}

// Code to execute if the user chooses to manage templates
func manageTemplatesMenuOption() {
	getUserSelection(
		"Choose a template:",
		getListOfTemplates(),
	)
}

// Code to execute if the user chooses to manage mixins
func manageMixinsMenuOption() {
	getUserSelection(
		"Choose a template:",
		getListOfTemplates(),
	)
}

// Takes in a list of possible choices and returns a map of numerical indexes to choices
func createOptionMap(options ...UserChoice) map[int]UserChoice {
	numberToOption := make(map[int]UserChoice)
	for i, option := range options {
		numberToOption[i+1 /* 0-index to 1-index */] = option
	}
	return numberToOption
}
