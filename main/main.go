package main

import (
	"fmt"
	"os"
)

const (
	BUILD_DOCKERFILE UserChoice = "1"
	MANAGE_TEMPLATES UserChoice = "2"
	MANAGE_MIXINS    UserChoice = "3"
	EXIT             UserChoice = "4"
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
	defaultSelection := "1"

	fmt.Printf(`Main Menu:
	1) build dockerfile from template
	2) manage templates
	3) manage mixins
	4) exit
What would you like to do? (%s): `, defaultSelection)

	userChoice := UserChoice(readLineFromStdInAsString(string(defaultSelection)))

	return userChoice
}
