package main

import (
	"fmt"
	"os"
	"strings"
)

var contentPath string = strings.Join([]string{home, ".config", "docker-composer", "mixins"}, string(os.PathSeparator))

func main() {
	dat, err := os.ReadFile("/home/zarinloosli/docker-composer/example.tplt")
	check(err)
	ast := tokenize(string(dat))
	for i, token := range ast {
		if token.kind == variable {
			ast[i] = handleVariable(token)
			fmt.Println(ast[i])
		}
	}
	// fmt.Print(ast)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
