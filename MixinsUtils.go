package main

import (
	"fmt"
	"os"
	"time"
)

var mixinDirPath = segmentsToPath(string(contentPath), "mixins")

func createMixin() (string, error) {
	createDirectoryIfDoesNotExist(mixinDirPath)

	now := time.Now().UnixNano()
	tempFilename := getMixinPathFromName(UserChoice(now))
	e := writeStringToFile("# your mixin here", tempFilename)
	if e != nil {
		panic(e)
	}
	defer os.Remove(string(tempFilename))

	// edit file
	fmt.Println("Opening mixin editor...")
	e = editFileInUserPreferredEditor(tempFilename)
	if e != nil {
		panic(e)
	}

	// save file with user-specified name
	for {
		fmt.Print("Choose a name for your mixin: ")
		userSpecifiedMixinName := readLineFromStdInAsString()

		newPath := getMixinPathFromName(userSpecifiedMixinName)

		// if name already in use, try again
		if fileExists(newPath) {
			fmt.Println("That name is already in use")
			continue
		}

		// if path good, proceed
		os.Rename(string(tempFilename), string(newPath))
		return getMixinContents(userSpecifiedMixinName)
	}
}

func getMixinPathFromName(name UserChoice) Path {
	return segmentsToPath(string(mixinDirPath), string(name))
}

func getMixinContents(name UserChoice) (string, error) {
	path := getMixinPathFromName(name)
	if fileExists(path) {
		return readStringFromFile(path)
	}
	return "", os.ErrNotExist
}

func getListOfMixins() []UserChoice {
	mixins, e := os.ReadDir(string(mixinDirPath))
	if e != nil {
		panic(e)
	}
	numberToMixin := []UserChoice{}

	for _, file := range mixins {
		filename := file.Name()
		numberToMixin = append(numberToMixin, UserChoice(filename))
	}

	return numberToMixin
}
