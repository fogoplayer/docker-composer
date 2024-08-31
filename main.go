package main

import (
	"fmt"
	"log/slog"
	"os"
	"path"
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
	userChoice, remainingArgs := handleCLArgs(os.Args)
	cliMode := len(remainingArgs) > 0

	for {
		if userChoice == INVALID {
			userChoice = getUserSelection(
				"What would you like to do?:",
				[]UserChoice{
					BUILD_DOCKERFILE,
					MANAGE_TEMPLATES,
					MANAGE_MIXINS,
					EXIT,
				},
			)
		}

		switch userChoice {
		case BUILD_DOCKERFILE:
			buildDockerfileMenuOption(remainingArgs...)

		case MANAGE_TEMPLATES:
			manageTemplatesMenuOption(remainingArgs...)

		case MANAGE_MIXINS:
			manageMixinsMenuOption(remainingArgs...)

		case EXIT:
			os.Exit(0)

		default:
			fmt.Println("Invalid input. Please try again")
			continue
		}

		if cliMode {
			os.Exit(0)
		} else {
			userChoice = INVALID
			remainingArgs = []string{}
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

	variables := make(map[string][]string)
	for i, token := range ast {
		if token.kind == VARIABLE {
			// don't read input if the name has been memoized
			if len((variables)[token.name]) == 0 {
				var values []string

				if len(defaults) > 0 {
					values = populateVariableWithMixins(token.name, UserChoice(defaults[0]))
					defaults = defaults[1:]
				} else {
					values = populateVariableWithMixins(token.name)
				}
				(variables)[token.name] = values
			}

			ast[i] = Token{values: (variables)[token.name]}
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

		templatePath := getTemplatePathFromName(selectedTemplate)
		if !fileExists(templatePath) {
			// Let the default block in the switch work for us
			selectedAction = INVALID
		}

		switch selectedAction {
		case EDIT:
			editFileInUserPreferredEditor(templatePath)
			break selectTemplateLoop

		case DELETE:
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
func manageMixinsMenuOption(args ...string) {
	const (
		CREATE_NEW UserChoice = "create new mixin"
		EDIT       UserChoice = "edit a mixin"
		DELETE     UserChoice = "delete a mixin"
		INVALID    UserChoice = "invalid"
	)

	var selectedAction UserChoice = INVALID
	var selectedMixin UserChoice = INVALID

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
		selectedMixin = UserChoice(args[0])
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
			createMixin()
			return
		}

		if selectedAction != CREATE_NEW && selectedAction != EDIT && selectedAction != DELETE {
			fmt.Println("Invalid input. Please try again")
			selectedAction = INVALID
			selectedMixin = INVALID
			continue selectActionLoop
		}
		break selectActionLoop
	}

selectTemplateLoop:
	for {
		if selectedMixin == INVALID {
			selectedMixin = getUserSelection(
				"Choose a mixin:",
				getListOfMixins(),
			)
		}

		mixinPath := getMixinPathFromName(selectedMixin)
		if !fileExists(mixinPath) {
			// Let the default block in the switch work for us
			selectedAction = INVALID
		}

		switch selectedAction {
		case EDIT:
			editFileInUserPreferredEditor(mixinPath)
			break selectTemplateLoop

		case DELETE:
			deleteFile(mixinPath)
			break selectTemplateLoop

		default:
			fmt.Println("Invalid input. Please try again")
			selectedAction = INVALID
			selectedMixin = INVALID
			continue selectTemplateLoop
		}
	}
}

// parses command-line arguments
func handleCLArgs(args []string) (UserChoice, []string) {
	workingDirectory = Path(path.Dir(args[0]))
	args = args[1:]

	if len(args) == 0 {
		return INVALID, []string{}
	}

	switch args[0] {
	case "dockerfile":
		return BUILD_DOCKERFILE, args[1:]

	case "template":
		return MANAGE_TEMPLATES, args[1:]

	case "mixin":
		return MANAGE_MIXINS, args[1:]

	default:
		return INVALID, []string{}
	}
}
