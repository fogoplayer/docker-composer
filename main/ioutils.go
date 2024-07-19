package main

import (
	"fmt"
	"os"
	"os/exec"
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
