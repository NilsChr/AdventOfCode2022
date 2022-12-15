package utils

import (
	"os"
	"runtime"
	"strings"
)

func GetInput(path string) []string {
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	content := string(dat)
	osName := runtime.GOOS

	endOfLine := "\n"
	if osName == "windows" {
		endOfLine = "\r\n"
	}

	var lines []string = strings.Split(string(content), endOfLine)

	return lines
}
