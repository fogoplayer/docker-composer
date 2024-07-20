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
var contentPath = segmentsToPath(home, ".config", "docker-composer")

// Entry point for docker-composer
func main() {
	for {
		userChoice := getUserSelection(
			"What would you like to do?:",
			[]UserChoice{
				BUILD_DOCKERFILE,
				MANAGE_TEMPLATES,
				MANAGE_MIXINS,
				EXIT,
			},
		)

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

// Code to execute if the user chooses to build a dockerfile.
//
// Reads in a template, tokenizes it, replaces variables, and saves to a directory of the user's choice
func buildDockerfileMenuOption() {
	const CREATE_NEW = "create a new template"
	templateList := getListOfTemplates()
	templateList[len(templateList)+1] = CREATE_NEW

	selectedTemplateName := getUserSelection("Choose a template:", templateList)

	var templateContents string
	var e error

	if selectedTemplateName == CREATE_NEW {
		templateContents, e = createTemplate()
		if e != nil {
			panic(e)
		}
	} else {
		templatePath := getTemplatePathFromName(selectedTemplateName)
		templateContents, e = readStringFromFile(templatePath)

		if e != nil {
			panic(e)
		}
	}

	ast := tokenize(templateContents)
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
			[]UserChoice{CREATE_NEW, EDIT, DELETE},
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
			templatePath := getTemplatePathFromName(selectedTemplate)
			editFileInUserPreferredEditor(templatePath)
			break manageTemplateLoop

		case DELETE:
			templatePath := getTemplatePathFromName(selectedTemplate)
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
			[]UserChoice{CREATE_NEW, EDIT, DELETE},
			"2",
		)

		if selectedAction == CREATE_NEW {
			createMixin()
			break manageMixinLoop
		}

		selectedMixin := getUserSelection(
			"Choose a mixin:",
			getListOfMixins(),
		)

		switch selectedAction {
		case EDIT:
			mixinPath := getMixinPathFromName(selectedMixin)
			editFileInUserPreferredEditor(mixinPath)
			break manageMixinLoop

		case DELETE:
			mixinPath := getMixinPathFromName(selectedMixin)
			deleteFile(mixinPath)
			break manageMixinLoop

		default:
			fmt.Println("Invalid input. Please try again")
			continue manageMixinLoop
		}
	}
}
