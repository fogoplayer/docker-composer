package main

import (
	"fmt"
	"os"
	"time"
)

var templateDirPath string = segmentsToPath(contentPath, "templates")

func createTemplate() string {
	// Create temporary template file
	os.MkdirAll(templateDirPath, os.ModePerm) // TODO slight vulnerability

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
	return segmentsToPath(templateDirPath, name)
}

func getTemplate(name string) string {
	template, _ := readStringFromFile(getTemplatePathFromName(name))
	return template
}

func getListOfTemplates() map[int]string {
	createBlankTemplateIfDoesNotExist()
	templates, _ := os.ReadDir(templateDirPath)
	numberToTemplate := make(map[int]string)

	for i, file := range templates {
		i = i + 1 // 0-index to 1-index
		filename := file.Name()
		numberToTemplate[i] = filename
	}

	return numberToTemplate
}

func createBlankTemplateIfDoesNotExist() {
	existing := getTemplate("blank")
	if existing == "" {
		os.MkdirAll(templateDirPath, os.ModePerm) // TODO slight vulnerability
		writeStringToFile(blankTemplate, getTemplatePathFromName("blank"))
	}
}

var blankTemplate string = `FROM {{base_image}}

{{contents}}`
