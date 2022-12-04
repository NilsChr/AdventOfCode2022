package main

import (
	"advent-of-code-2022/day1"
	"advent-of-code-2022/day2"
	"advent-of-code-2022/day3"
	"advent-of-code-2022/day4"
	"fmt"
	"os"
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
		fmt.Println("Running day1")
		day1.Day1()
	case "2":
		fmt.Println("Running day2")
		day2.Day2()
	case "3":
		fmt.Println("Running day3")
		day3.Day3()
	case "4":
		fmt.Println("Running day4")
		day4.Day4()
	}
}
