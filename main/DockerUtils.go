package main

import (
	"fmt"
	"strings"
)

func buildDockerfileFromAst(ast []Token) string {
	var dockerfileBuilder strings.Builder

	for _, token := range ast {
		for _, value := range token.values {
			fmt.Fprint(&dockerfileBuilder, value)
		}
	}

	return dockerfileBuilder.String()
}

func saveDockerFile(dockerfile string) {
	for {
		fmt.Print("Enter a directory to save your dockerfile to (./): ")
		savePath := segmentsToPath(string(
			readLineFromStdInAsString(UserChoice(workingDirectory)),
		))
		// make sure directory exists
		if !fileExists(savePath) {
			fmt.Println("There is an issue with that directory. Please try again.")
			continue
		}

		savePath = segmentsToPath(string(savePath), "dockerfile")
		writeStringToFile(dockerfile, savePath)

		fmt.Printf("dockerfile saved to %s\n", savePath)
		break
	}
}
