package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var mixinDirPath string = segmentsToPath(contentPath, "mixins")

func createMixin() string {
	// Create temporary mixin file
	os.MkdirAll(mixinDirPath, os.ModePerm) // TODO slight vulnerability

	now := time.Now().UnixNano()
	tempFilename := getMixinPathFromName(string(now))
	os.WriteFile(tempFilename, []byte("# your mixin here"), os.ModePerm)
	// TODO cleanup with defer in case of error

	// edit file
	fmt.Println("Opening mixin editor...")
	openFileInUserPreferredEditor(tempFilename)

	// save file with user-specified name
	var userSpecifiedMixinName string
	fmt.Print("Choose a name for your mixin: ")
	fmt.Scanln(&userSpecifiedMixinName)

	os.Rename(tempFilename, getMixinPathFromName(userSpecifiedMixinName))

	return getMixin(userSpecifiedMixinName)
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
	return segmentsToPath(mixinDirPath, name+".mxin")
}

func getMixin(name string) string {
	filePath := getMixinPathFromName(name)
	bytes, _ := os.ReadFile(filePath)
	return string(bytes)
}

func listMixins() map[int]string {
	mixins, _ := os.ReadDir(mixinDirPath)
	numberToMixin := make(map[int]string)

	for i, file := range mixins {
		i = i + 1 // 0-index to 1-index
		filename := file.Name()
		filename = filename[:len(filename)-5]

		fmt.Println("\t", i, ")", filename)
		numberToMixin[i] = filename
	}

	return numberToMixin
}
