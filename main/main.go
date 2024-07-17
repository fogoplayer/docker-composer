package main

import (
	"fmt"
	"os"
)

func main() {
	dat, err := os.ReadFile("/home/zarinloosli/docker-composer/dockerfile-kasm-ubuntu-jammy-desktop.txt")
	check(err)
	ast := tokenize(string(dat))
	fmt.Print(ast)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
