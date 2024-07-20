package main

import (
	"fmt"
	"os"
	"time"
)

var mixinDirPath string = segmentsToPath(contentPath, "mixins")

func createMixin() (string, error) {
	// Create temporary mixin file
	os.MkdirAll(mixinDirPath, os.ModePerm) // TODO slight vulnerability

	now := time.Now().UnixNano()
	tempFilename := getMixinPathFromName(UserChoice(now))
	writeStringToFile("# your mixin here", tempFilename)
	// TODO cleanup with defer in case of error

	// edit file
	fmt.Println("Opening mixin editor...")
	openFileInUserPreferredEditor(tempFilename)

	// save file with user-specified name
	fmt.Print("Choose a name for your mixin: ")
	userSpecifiedMixinName := readLineFromStdInAsString()

	os.Rename(tempFilename, getMixinPathFromName(userSpecifiedMixinName))

	return getMixin(userSpecifiedMixinName)
}

func getMixinPathFromName(name UserChoice) string {
	return segmentsToPath(mixinDirPath, string(name)+".mxin")
}

func getMixin(name UserChoice) (string, error) {
	return readStringFromFile(getMixinPathFromName(name))
}

func getListOfMixins() map[int]UserChoice {
	mixins, _ := os.ReadDir(mixinDirPath)
	numberToMixin := make(map[int]UserChoice)

	for i, file := range mixins {
		i = i + 1 // 0-index to 1-index
		filename := file.Name()
		filename = filename[:len(filename)-5]
		numberToMixin[i] = UserChoice(filename)
	}

	return numberToMixin
}
