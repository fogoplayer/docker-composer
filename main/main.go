package main

import (
	"fmt"
	"os"
	"strings"
)

var home string = os.Getenv("HOME")
var contentPath string = segmentsToPath(home, ".config", "docker-composer")

func main() {
	template := readStringFromFile("/home/zarinloosli/docker-composer/example.tplt")
	ast := tokenize(template)
	for i, token := range ast {
		if token.kind == variable {
			ast[i] = handleVariable(token)
		}
	}

	dockerfile := buildDockerfile(ast)
	saveDockerFile(dockerfile)
}
