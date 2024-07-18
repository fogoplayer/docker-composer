package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("/home/zarinloosli/docker-composer/example.tplt")
	check(err)
	ast := tokenize(string(dat))
	for i, token := range ast {
		ast[i] = handleVariable(token)
	}
	fmt.Print(ast)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
