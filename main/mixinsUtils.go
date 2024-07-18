package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var home string = os.Getenv("HOME")
var mixinDirPath string = strings.Join([]string{home, ".config", "docker-composer"}, string(os.PathSeparator))

func createMixin() {
	// Create temporary mixin file
	os.MkdirAll(mixinDirPath, os.ModePerm) // TODO slight vulnerability

	now := time.Now().UnixNano()
	tempFilename := getMixinPathFromName(string(now))
	os.WriteFile(tempFilename, []byte("your mixin here"), os.ModePerm)

	// edit file
	fmt.Println("Opening mixin editor...")
	openFileInUserPreferredEditor(tempFilename)

	// save file with user-specified name
	var userSpecifiedMixinName string
	fmt.Print("Choose a name for your mixin: ")
	fmt.Scanln(&userSpecifiedMixinName)

	os.Rename(tempFilename, getMixinPathFromName(userSpecifiedMixinName))
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

func getMixinPathFromName(name string) string {
	return mixinDirPath + string(os.PathSeparator) + name + ".mxin"
}
