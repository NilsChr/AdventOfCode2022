package utils

import (
	"os"
	"strings"
)

func GetInput(path string) []string {
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	content := string(dat)
	var lines []string = strings.Split(string(content), "\n")
	/*for i, line := range lines {
		lines[i] = line[:len(line)-1]
	}*/
	return lines
}
