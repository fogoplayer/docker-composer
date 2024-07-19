package main

import (
	"os"
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
