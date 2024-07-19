package main

import (
	"fmt"
	"os"
	"strings"
)

var home string = os.Getenv("HOME")
var contentPath string = segmentsToPath(home, ".config", "docker-composer")

func main() {
	dat, err := os.ReadFile("/home/zarinloosli/docker-composer/example.tplt")
	check(err)
	ast := tokenize(string(dat))
	for i, token := range ast {
		if token.kind == variable {
			ast[i] = handleVariable(token)
		}
	}

	dockerfile := buildDockerfile(ast)
	saveDockerFile(dockerfile)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func buildDockerfile(ast []Token) string {
	var dockerfileBuilder strings.Builder

	for _, token := range ast {
		for _, value := range token.values {
			fmt.Fprint(&dockerfileBuilder, value)
		}
	}

	return dockerfileBuilder.String()
}

func saveDockerFile(dockerfile string) {
	var savePath string
	fmt.Print("Enter a path to save your dockerfile to (~/): ")
	fmt.Scanln(&savePath)
	if savePath == "" {
		savePath = os.Getenv("HOME")
	}

	savePath = segmentsToPath(savePath, "dockerfile")
	os.WriteFile(savePath, []byte(dockerfile), os.ModePerm)
	fmt.Printf("dockerfile saved to %s\n", savePath)
}
