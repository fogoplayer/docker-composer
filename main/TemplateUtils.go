package main

import (
	"fmt"
	"os"
	"time"
)

var templateDirPath = segmentsToPath(string(contentPath), "templates")

// Creates a new template
//
// Opens the user's preferred text editor, lets them add template contents, select a name, and save
func createTemplate() (string, error) {
	// Create temporary template file
	os.MkdirAll(string(templateDirPath), os.ModePerm) // TODO slight vulnerability

	now := time.Now().UnixNano()
	tempFilename := getTemplatePathFromName(UserChoice(now))
	writeStringToFile("# your template here", tempFilename)
	// TODO cleanup with defer in case of error

	// edit file
	fmt.Println("Opening template editor...")
	editFileInUserPreferredEditor(tempFilename)

	// save file with user-specified name
	fmt.Print("Choose a name for your template: ")
	userSpecifiedTemplateName := readLineFromStdInAsString()

	path := getTemplatePathFromName(userSpecifiedTemplateName)
	os.Rename(string(tempFilename), string(path))

	return getTemplateContents(userSpecifiedTemplateName)
}

// Takes in a template name and returns the path
func getTemplatePathFromName(name UserChoice) Path {
	return segmentsToPath(string(templateDirPath), string(name))
}

// Takes in a template name and returns the contents
func getTemplateContents(name UserChoice) (string, error) {
	path := getTemplatePathFromName(name)
	return readStringFromFile(path)

}

// Returns a map of numerical indexes to names of all created templates
func getListOfTemplates() []UserChoice {
	createBlankTemplateIfDoesNotExist()
	templates, e := os.ReadDir(string(templateDirPath))
	if e != nil {
		panic(e)
	}
	templateList := []UserChoice{}

	for _, file := range templates {
		filename := file.Name()
		templateList = append(templateList, UserChoice(filename))
	}

	return templateList
}

// Checks if the `blank` template exists. If not, creates one using a default template hard-coded into the source code.
func createBlankTemplateIfDoesNotExist() error {
	_, e := getTemplateContents("blank")
	if e == nil {
		return nil
	}

	if errorIsNotThatFileExists(e) {
		panic(e)
	}

	createDirectoryIfDoesNotExist(templateDirPath)
	path := getTemplatePathFromName("blank")
	return writeStringToFile(blankTemplate, path)
}

var blankTemplate string = `FROM {{base_image}}

{{contents}}`
