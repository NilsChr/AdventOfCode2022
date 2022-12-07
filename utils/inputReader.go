package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GetInput(path string) []string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error", err)
		panic("Could not read file")
	}

	var lines []string = strings.Split(string(content), "\n")

	return lines
}

func GetInput2(path string) []string {
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	content := string(dat)
	var lines []string = strings.Split(string(content), "\n")
	for i, line := range lines {
		lines[i] = line[:len(line)-1]
	}
	return lines
}
