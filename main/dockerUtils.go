package main

import (
	"fmt"
	"os"
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
	fmt.Print("Enter a path to save your dockerfile to (~/): ")
	savePath := Path(readLineFromStdInAsString(UserChoice(os.Getenv("HOME"))))

	savePath = segmentsToPath(string(savePath), "dockerfile")
	writeStringToFile(dockerfile, savePath)
	fmt.Printf("dockerfile saved to %s\n", savePath)
}
