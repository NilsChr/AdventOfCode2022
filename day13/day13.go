package day13

import (
	"advent-of-code-2022/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day13() {
	lines := utils.GetInput("./day13/input-test.txt")
	fmt.Println("Task1: ", task1(lines))
}

func task1(lines []string) int {

	/*for _, line := range lines {

		consumeLine(line)

	}
	*/
	pairs := 0
	//rightOrder := 0
	sums := 0
	for i := 0; i < len(lines); i += 3{
		pairs++

		a := lines[i]
		b := lines[i+1]

		aSlice := stringToSlice(reduceTo1D(a))
		bSlice := stringToSlice(reduceTo1D(b))

		fmt.Println(aSlice)
		fmt.Println(bSlice)
		fmt.Println()

		order := checkSliceOrder(aSlice, bSlice)

		if order == -1 {
			fmt.Println("Correct:",pairs)
			sums += pairs
		}

	}

	return sums
}

func reduceTo1D(line string) string {
	for checkDepth(line) > 1 {
		output, _ := consumeLine2(line)
		//fmt.Println(output, consumedSlice)
		line = output
	}
	return line
}

func consumeLine(line string) (string, []int){
	startBracket := 0
	endBracket := 0

	for i := 0; i < len(line); i++ {
		part := string(line[i])
		if part == "[" {
			startBracket = i
		} 
		if part == "]" {
			endBracket = i
			break;
		}
	}

	if startBracket == 0 && endBracket == len(line)-1 {
		return "", stringToSlice(line)
	}

	slice := line[startBracket:endBracket+1]
	line = strings.Replace(line, slice, "",1)

	if line[startBracket] == ',' {
		line = line[0:startBracket] + line[startBracket+1:]
	}

	return line, stringToSlice(slice)
}

func consumeLine2(line string) (string, []int){
	startBracket := 0
	endBracket := 0

	for i := 0; i < len(line); i++ {
		part := string(line[i])
		if part == "[" {
			startBracket = i
		} 
		if part == "]" {
			endBracket = i
			break;
		}
	}

	if startBracket == 0 && endBracket == len(line)-1 {
		return "", stringToSlice(line)
	}

	slice := line[startBracket:endBracket+1]

	sliceData := stringToSlice(slice)
	val := collapseSlice(sliceData)
	/*
	fmt.Println("Got val", val, " from slice: ", sliceData)
	fmt.Println("Cutting out: ", slice, " replace with ", val, "=", strconv.Itoa(val))
	fmt.Println("Line before:",line)
	*/
	line = strings.Replace(line, slice,strconv.Itoa(val),1)
	//fmt.Println("Line after :",line)

	if line[startBracket] == ',' {
		line = line[0:startBracket] + line[startBracket+1:]
	}

	return line, stringToSlice(slice)
}

func stringToSlice(line string) []int {
	var out []int

	if line == "[]" {
		out = append(out, -1)
		return out
	}

	//fmt.Println("In: ", line)
	line = strings.Replace(line,"[", "",1)
	line = strings.Replace(line,"]", "",1)


	numbers := strings.Split(line, ",")

//	fmt.Println("numbers", numbers, len(numbers))
	
	for _,n := range numbers {
		parsed,_ := strconv.Atoi(n)
		out = append(out, parsed)
	}

	return out
}

func collapseSlice(slice []int) int {
	if len(slice) == 0 {
		return -1
	}

	min, max := MinMax(slice)
	//fmt.Println("Min, max", min, max)
	if sort.IntsAreSorted(slice) {
		return max
	}
	return min
}

func checkDepth(line string) int {
	open := strings.Count(line, "[")
	//close := strings.Count(line, "]")
	return open
}

func MinMax(array []int) (int, int) {
    var max int = array[0]
    var min int = array[0]
    for _, value := range array {
        if max < value {
            max = value
        }
        if min > value {
            min = value
        }
    }
    return min, max
}

func checkSliceOrder(a []int, b []int) int {

	if len(a) < len(b) {
		return 1
	}

	if len(b) < len(a) {
		return -1
	}

	for j := 0; j < len(a); j++ {
		if a[j] < b[j] {
			return -1
		}
		if b[j] < a[j] {
			//rightOrder++
			//sums += pairs
			//fmt.Println("not in right order", pairs)
			return 1
		}
	}
	return 0
}

