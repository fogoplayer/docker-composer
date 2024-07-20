package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Path string

//////////////
// File I/O //
//////////////

func segmentsToPath(segments ...string) Path {
	return Path(strings.Join(segments, string(os.PathSeparator)))
}

func writeStringToFile(data string, path Path) error {
	return os.WriteFile(string(path), []byte(data), os.ModePerm)
}

func readStringFromFile(path Path) (string, error) {
	data, e := os.ReadFile(string(path))

	return string(data), e
}

func editFileInUserPreferredEditor(filename Path) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = os.Getenv("VISUAL")
	}
	if editor == "" {
		editor = "/usr/bin/editor"
	}

	editorProcess := exec.Command(editor, string(filename))
	editorProcess.Stdin = os.Stdin
	editorProcess.Stdout = os.Stdout
	editorProcess.Stderr = os.Stderr
	return editorProcess.Run()
}

func deleteFile(path Path) {
	os.Remove(string(path))
}

func createDirectoryIfDoesNotExist(path Path) error {
	return os.MkdirAll(string(path), os.ModePerm) // TODO slight vulnerability
}

func fileExists(path Path) bool {
	_, e := os.Stat(string(path))
	return e == nil
}

func errorIsNotThatFileExists(e error) bool {
	return e == nil || !errors.Is(e, os.ErrExist) || errors.Is(e, os.ErrNotExist)
}

///////////////////
// STDIN/OUT I/O //
///////////////////

func readLineFromStdInAsString(defaultValue ...UserChoice) UserChoice {
	var userInput UserChoice
	fmt.Scanln(&userInput)
	if len(userInput) > 0 || len(defaultValue) < 1 {
		return userInput
	} else {
		return defaultValue[0]
	}
}

func printSelectionList(mapObj map[int]UserChoice) {
	for i := 1; i <= len(mapObj); i++ {
		val := mapObj[i]
		fmt.Println("  "+strconv.Itoa(i)+")", val)
	}
}

func getUserSelection(message string, numberToOption map[int]UserChoice, defaultValue ...UserChoice) UserChoice {
	if len(defaultValue) == 0 {
		defaultValue = []UserChoice{"1"}
	}

	fmt.Println(message)
	printSelectionList(numberToOption)
	fmt.Printf("Choose an option (%s): ", defaultValue[0])
	userChoice := readLineFromStdInAsString(defaultValue[0])

	// if default
	if userChoice == "" {
		userChoice = defaultValue[0]
	}

	// if a number
	num, e := strconv.Atoi(string(userChoice))
	if e == nil {
		return numberToOption[num]
	}

	// if a string
	return userChoice
}
