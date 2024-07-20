package main

import (
	"bufio"
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
	stdin := bufio.NewReader(os.Stdin)
	userInput, e := stdin.ReadString('\n') // read until newline
	if e != nil {
		panic(e)
	}

	if len(userInput) > 0 || len(defaultValue) < 1 {
		return UserChoice(userInput)
	} else {
		return defaultValue[0]
	}
}

func printSelectionList(choices []UserChoice) {
	for i := 0; i < len(choices); i++ {
		val := choices[i]
		fmt.Println("  "+strconv.Itoa(i+1)+")", val)
	}
}

func getUserSelection(message string, numberToOption []UserChoice, defaultValue ...UserChoice) UserChoice {
	if len(defaultValue) == 0 {
		defaultValue = []UserChoice{"1"}
	}

	fmt.Println(message)
	printSelectionList(numberToOption)
	fmt.Printf("Choose an option (%s): ", defaultValue[0])
	userChoice := readLineFromStdInAsString(defaultValue[0])
	userChoice = UserChoice(strings.Replace(string(userChoice), "\n", "", -1))

	// if default
	if userChoice == "" {
		userChoice = defaultValue[0]
	}

	// if a number
	num, e := strconv.Atoi(string(userChoice))
	if e == nil && num > 0 {
		return numberToOption[num-1]
	}

	// if a string
	return userChoice
}
