package main

import (
	"advent-of-code-2022/day1"
	"advent-of-code-2022/day10"
	"advent-of-code-2022/day2"
	"advent-of-code-2022/day3"
	"advent-of-code-2022/day4"
	"advent-of-code-2022/day5"
	"advent-of-code-2022/day6"
	"advent-of-code-2022/day7"
	"advent-of-code-2022/day8"
	"advent-of-code-2022/day9"
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("no arguments added")
		return
	}
	day := os.Args[1:][0]
	start := time.Now()
	
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
	case "5":
		fmt.Println("Running day5")
		day5.Day5()
	case "6":
		fmt.Println("Running day6")
		day6.Day6()
	case "7":
		fmt.Println("Running day7")
		day7.Day7()
	case "8":
		fmt.Println("Running day8")
		day8.Day8()
	case "9":
		fmt.Println("Running day8")
		day9.Day9()
	case "10":
		fmt.Println("Running day10")
		day10.Day10()
	}


	elapsed := time.Since(start)
	fmt.Printf("Time %f seconds\n", elapsed.Seconds())

}
