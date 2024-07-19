package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var templateDirPath string = segmentsToPath(contentPath, "templates")

func createTemplate() string {
	// Create temporary mixin file
	os.MkdirAll(mixinDirPath, os.ModePerm) // TODO slight vulnerability

	now := time.Now().UnixNano()
	tempFilename := getTemplatePathFromName(string(now))
	writeStringToFile("# your template here", tempFilename)
	// TODO cleanup with defer in case of error

	// edit file
	fmt.Println("Opening template editor...")
	openFileInUserPreferredEditor(tempFilename)

	// save file with user-specified name
	fmt.Print("Choose a name for your template: ")
	userSpecifiedTemplateName := readLineFromStdInAsString()

	os.Rename(tempFilename, getTemplatePathFromName(userSpecifiedTemplateName))

	return getTemplate(userSpecifiedTemplateName)
}

func getTemplatePathFromName(name string) string {
	return segmentsToPath(mixinDirPath, name+".mxin")
}

func getTemplate(name string) string {
	return readStringFromFile(getTemplatePathFromName(name))
}

func listTemplates() map[int]string {
	templates, _ := os.ReadDir(templateDirPath)
	numberToTemplate := make(map[int]string)

	for i, file := range templates {
		i = i + 2 // 0-index to 1-index, with offset
		filename := file.Name()
		filename = filename[:len(filename)-5]

		fmt.Println("\t", strconv.Itoa(i)+")", filename)
		numberToTemplate[i] = filename
	}

	return numberToTemplate
}

var defaultTemplate string = `FROM {{base_image}}
USER root

{{setup}}`
