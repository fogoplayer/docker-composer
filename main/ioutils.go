package main

import (
	"os"
	"strings"
)

func segmentsToPath(segments ...string) string {
	return strings.Join(segments, string(os.PathSeparator))
}
