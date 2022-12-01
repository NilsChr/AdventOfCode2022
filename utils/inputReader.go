package utils

import (
	"fmt"
	"io/ioutil"
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