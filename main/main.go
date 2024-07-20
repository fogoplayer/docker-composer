package main

import (
	"fmt"
	"log/slog"
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
	os.Args = []string{"template", "delete", "testing\n"}
	handleCLArgs(os.Args)

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
// # Reads in a template, tokenizes it, replaces variables, and saves to a directory of the user's choice
//
// Optionally, strings can be passed in to bypass selections. The first string is the name of a template, and the following
// strings will be passed into each variable slot 1:1.
func buildDockerfileMenuOption(defaults ...string) {
	const CREATE_NEW = "create a new template"
	var selectedTemplateName UserChoice
	var templateContents string

	// choose a template
	if len(defaults) > 0 {
		selectedTemplateName = UserChoice(defaults[0])
		defaults = defaults[1:]
	} else {
		selectedTemplateName = getUserSelection(
			"Choose a template:",
			append(
				getListOfTemplates(),
				CREATE_NEW,
			),
		)
	}

	// get template contents
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
			if len(defaults) > 0 {
				ast[i] = populateVariableWithMixins(token, UserChoice(defaults[0]))
				defaults = defaults[1:]
			} else {
				ast[i] = populateVariableWithMixins(token)
			}
		}
	}

	if len(defaults) > 0 {
		slog.Warn("Some passed-in values were unused")
	}

	dockerfile := buildDockerfileFromAst(ast)
	saveDockerFile(dockerfile)
}

// Code to execute if the user chooses to manage templates
func manageTemplatesMenuOption(args ...string) {
	const (
		CREATE_NEW UserChoice = "create new template"
		EDIT       UserChoice = "edit a template"
		DELETE     UserChoice = "delete a template"
		INVALID    UserChoice = "invalid"
	)

	var selectedAction UserChoice = INVALID
	var selectedTemplate UserChoice = INVALID

	if len(args) > 0 {
		switch args[0] {
		case "create":
			selectedAction = CREATE_NEW
		case "edit":
			selectedAction = EDIT
		case "delete":
			selectedAction = DELETE
		}
		args = args[1:]
	}

	if len(args) > 0 {
		selectedTemplate = UserChoice(args[0])
	}

selectActionLoop:
	for {
		if selectedAction == INVALID {
			selectedAction = getUserSelection(
				"What action would you like to perform?",
				[]UserChoice{CREATE_NEW, EDIT, DELETE},
				"2",
			)
		}

		if selectedAction == CREATE_NEW {
			createTemplate()
			return
		}

		if selectedAction != CREATE_NEW && selectedAction != EDIT && selectedAction != DELETE {
			fmt.Println("Invalid input. Please try again")
			selectedAction = INVALID
			selectedTemplate = INVALID
			continue selectActionLoop
		}
		break selectActionLoop
	}

selectTemplateLoop:
	for {
		if selectedTemplate == INVALID {
			selectedTemplate = getUserSelection(
				"Choose a template:",
				getListOfTemplates(),
			)
		}

		switch selectedAction {
		case EDIT:
			templatePath := getTemplatePathFromName(selectedTemplate)
			editFileInUserPreferredEditor(templatePath)
			break selectTemplateLoop

		case DELETE:
			templatePath := getTemplatePathFromName(selectedTemplate)
			deleteFile(templatePath)
			break selectTemplateLoop

		default:
			fmt.Println("Invalid input. Please try again")
			selectedAction = INVALID
			selectedTemplate = INVALID
			continue selectTemplateLoop
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

// parses command-line arguments
func handleCLArgs(args []string) {
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "dockerfile":
		buildDockerfileMenuOption(args[1:]...)
		os.Exit(0)

	case "template":
		manageTemplatesMenuOption(args[1:]...)
		os.Exit(0)

	case "mixin":
		manageMixinsMenuOption()
		os.Exit(0)

	default:
		fmt.Println("Unrecognized parameter")
	}
}
