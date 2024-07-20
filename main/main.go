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
var contentPath string = segmentsToPath(home, ".config", "docker-composer")

func main() {
	for {
		userChoice := getUserMainMenuChoice()

		switch userChoice {
		case BUILD_DOCKERFILE:
			template, _ := readStringFromFile("/home/zarinloosli/docker-composer/example.tplt")
			ast := tokenize(template)
			for i, token := range ast {
				if token.kind == variable {
					ast[i] = handleVariable(token)
				}
			}

			dockerfile := buildDockerfile(ast)
			saveDockerFile(dockerfile)

		case MANAGE_TEMPLATES:
			printSelectionList(getListOfTemplates())
			println(2)

		case MANAGE_MIXINS:
			println(3)

		case EXIT:
			os.Exit(0)

		default:
			fmt.Println("Invalid input. Please try again")
		}
	}
}

func getUserMainMenuChoice() UserChoice {
	menuOptions := createOptionMap(
		BUILD_DOCKERFILE,
		MANAGE_TEMPLATES,
		MANAGE_MIXINS,
		EXIT,
	)
	userChoice := getUserSelection("What would you like to do?:", menuOptions, BUILD_DOCKERFILE)

	return userChoice
}
