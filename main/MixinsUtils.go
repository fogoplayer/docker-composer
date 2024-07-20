package main

import (
	"fmt"
	"os"
	"time"
)

var mixinDirPath Path = segmentsToPath(string(contentPath), "mixins")

func createMixin() (string, error) {
	createDirectoryIfDoesNotExist(mixinDirPath)

	now := time.Now().UnixNano()
	tempFilename := getMixinPathFromName(UserChoice(now))
	writeStringToFile("# your mixin here", tempFilename)
	// TODO cleanup with defer in case of error

	// edit file
	fmt.Println("Opening mixin editor...")
	editFileInUserPreferredEditor(tempFilename)

	// save file with user-specified name
	fmt.Print("Choose a name for your mixin: ")
	userSpecifiedMixinName := readLineFromStdInAsString()

	os.Rename(string(tempFilename), string(getMixinPathFromName(userSpecifiedMixinName)))

	return getMixinContents(userSpecifiedMixinName)
}

func getMixinPathFromName(name UserChoice) Path {
	return segmentsToPath(string(mixinDirPath), string(name))
}

func getMixinContents(name UserChoice) (string, error) {
	return readStringFromFile(getMixinPathFromName(name))
}

func getListOfMixins() map[int]UserChoice {
	mixins, _ := os.ReadDir(string(mixinDirPath))
	numberToMixin := make(map[int]UserChoice)

	for i, file := range mixins {
		i = i + 1 // 0-index to 1-index
		filename := file.Name()
		numberToMixin[i] = UserChoice(filename)
	}

	return numberToMixin
}
