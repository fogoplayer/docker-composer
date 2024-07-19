package main

import (
	"fmt"
	"os"
	"strings"
)

func segmentsToPath(segments ...string) string {
	return strings.Join(segments, string(os.PathSeparator))
}

func readStringFromFile(path string) string {
	data, _ := os.ReadFile(path)
	return string(data)
}

func writeStringToFile(data string, path string) {
	os.WriteFile(path, []byte(data), os.ModePerm)
}

func readLineFromStdInAsString(defaultValue ...string) string {
	var userInput string
	fmt.Scanln(&userInput)
	if len(userInput) > 0 || len(defaultValue) < 1 {
		return userInput
	} else {
		return defaultValue[0]
	}
}
