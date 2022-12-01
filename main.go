package main

import (
	"os"
	"fmt"
	"advent-of-code-2022/day1"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("no arguments added")
		return
	}
	day := os.Args[1:][0]

	fmt.Println("advent ", day)
	switch day {
	case "1":
		fmt.Println("Running day1");
		day1.Day1()
	case "2":
		fmt.Println("Running day2"); 
	}
}