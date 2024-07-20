package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func segmentsToPath(segments ...string) string {
	return strings.Join(segments, string(os.PathSeparator))
}

func readStringFromFile(path string) (string, error) {
	data, err := os.ReadFile(path)

	return string(data), err
}

func writeStringToFile(data string, path string) {
	os.WriteFile(path, []byte(data), os.ModePerm)
}

func readLineFromStdInAsString(defaultValue ...UserChoice) UserChoice {
	var userInput UserChoice
	fmt.Scanln(&userInput)
	if len(userInput) > 0 || len(defaultValue) < 1 {
		return userInput
	} else {
		return defaultValue[0]
	}
}

func openFileInUserPreferredEditor(filename string) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = os.Getenv("VISUAL")
	}
	if editor == "" {
		editor = "/usr/bin/editor"
	}

	editorProcess := exec.Command(editor, filename)
	editorProcess.Stdin = os.Stdin
	editorProcess.Stdout = os.Stdout
	editorProcess.Stderr = os.Stderr
	editorProcess.Run()
}

func printSelectionList(mapObj map[int]UserChoice) {
	for i, val := range mapObj {
		fmt.Println("\t", strconv.Itoa(i)+")", val)
	}
}

func getUserSelection(message string, numberToOption map[int]UserChoice, defaultValue ...UserChoice) UserChoice {
	fmt.Println(message)
	printSelectionList(numberToOption)
	fmt.Printf("Selection (%s): ", defaultValue[0])
	userChoice := readLineFromStdInAsString(defaultValue[0])

	// if a number
	num, err := strconv.Atoi(string(userChoice))
	if err == nil {
		return numberToOption[num]
	}

	// if a string
	return userChoice
}

func createOptionMap(options ...string) map[int]UserChoice {
	numberToOption := make(map[int]UserChoice)
	for i, option := range options {
		numberToOption[i+1 /* 0-index to 1-index */] = UserChoice(option)
	}
	return numberToOption
}
