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
var contentPath, _ = segmentsToPath(home, ".config", "docker-composer")

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
	const (
		CREATE_NEW UserChoice = "create new template"
		EDIT       UserChoice = "edit a template"
		DELETE     UserChoice = "delete a template"
	)

manageTemplateLoop:
	for {
		selectedAction := getUserSelection(
			"What action would you like to perform?",
			createOptionMap(CREATE_NEW, EDIT, DELETE),
			"2",
		)

		if selectedAction == CREATE_NEW {
			createTemplate()
			break
		}

		selectedTemplate := getUserSelection(
			"Choose a template:",
			getListOfTemplates(),
		)

		switch selectedAction {
		case EDIT:
			templatePath, _ := getTemplatePathFromName(selectedTemplate)
			editFileInUserPreferredEditor(templatePath)
			break manageTemplateLoop

		case DELETE:
			templatePath, _ := getTemplatePathFromName(selectedTemplate)
			deleteFile(templatePath)
			break manageTemplateLoop

		default:
			fmt.Println("Invalid input. Please try again")
			continue
		}
	}
}

// Code to execute if the user chooses to manage mixins
func manageMixinsMenuOption() {
	const (
		CREATE_NEW UserChoice = "create new mixin"
		EDIT       UserChoice = "edit a mixin"
		DELETE     UserChoice = "delete a mixin"
	)
manageMixinLoop:
	for {
		selectedAction := getUserSelection(
			"What action would you like to perform?",
			createOptionMap(CREATE_NEW, EDIT, DELETE),
			"2",
		)

		if selectedAction == CREATE_NEW {
			createTemplate()
			break
		}

		selectedMixin := getUserSelection(
			"Choose a mixin:",
			getListOfMixins(),
		)

		switch selectedAction {
		case EDIT:
			mixinPath, _ := getMixinPathFromName(selectedMixin)
			editFileInUserPreferredEditor(mixinPath)
			break manageMixinLoop

		case DELETE:
			mixinPath, _ := getMixinPathFromName(selectedMixin)
			deleteFile(mixinPath)
			break manageMixinLoop

		default:
			fmt.Println("Invalid input. Please try again")
			continue
		}
	}
}

// Takes in a list of possible choices and returns a map of numerical indexes to choices
// TODO that's called a list... so just return a list
func createOptionMap(options ...UserChoice) map[int]UserChoice {
	numberToOption := make(map[int]UserChoice)
	for i, option := range options {
		numberToOption[i+1 /* 0-index to 1-index */] = option
	}
	return numberToOption
}
