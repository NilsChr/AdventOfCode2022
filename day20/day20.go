package day20

import (
	"advent-of-code-2022/utils"
	"fmt"
	"math"
	"strconv"
)

func Day20() {
	lines := utils.GetInput("./day20/input.txt")
	fmt.Println("Task1:", task1(lines))
/*
	test1 := []int{1, 2, -3, 3, -2, 0, 4}
	exp1  := []int{2, 1, -3, 3, -2, 0, 4}
	test1 = moveValue(test1, 0)
	fmt.Println("test 1 PASS:", check(test1, exp1))

	test2 := []int{4, -2, 5, 6, 7, 8, 9}
	exp2  := []int{4, 5, 6, 7, 8, -2, 9}
	test2 = moveValue(test2, 1)
	fmt.Println("test 2 PASS:", check(test2, exp2))
	//fmt.Println(test2, exp2)

	test3 := []int{1, -3, 2, 3, -2, 0, 4}
	exp3  := []int{1, 2, 3, -2, -3, 0, 4}
	test3 = moveValue(test3, 1)
	fmt.Println("test 3 PASS:", check(test3, exp3))

	test4 := []int{1, -3, 2, 3, -2, 0, 4}
	exp4  := []int{1, -3, 2, 3, -2, 0, 4}
	test4 = moveValue(test4, 5)
	fmt.Println("test 4 PASS:", check(test4, exp4))

	test5 := []int{1, 2, -2, -3, 0, 3, 4}
	exp5  := []int{1, 2, -3, 0, 3, 4, -2}
	test5 = moveValue(test5, 2)
	fmt.Println("test 5 PASS:", check(test5, exp5))
	fmt.Println(test5, exp5)

	test6 := []int{1, 2, -3, 0, 3, 4, -2}
	exp6  := []int{1, 2, -3, 4, 0, 3, -2}
	test6 = moveValue(test6, 5)
	fmt.Println("test 6 PASS:", check(test6, exp6))
*/
/*
	test7 := []int{1, 2, -3, 0, 3, 1, -2}
	exp7  := []int{1, 2, -3, 0, 3, -2, 1}
	test7 = moveValue(test7, 5)
	fmt.Println("test 7 PASS:", check(test7, exp7))
	fmt.Println(test7, exp7)
	*/
}

// too high : 16295
// too high : 7851
// Not right: -9344
// Not right: -9792
// not right: -5759
func task1(lines []string) int {
	file := parseLines(lines)
	queue := make([]int, len(file))
	copy(queue, file)
	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		nextIndex := -1
		for i := 0; i < len(file); i++ {
			if file[i] == next {
				nextIndex = i
				break
			}
		}
		file = moveValue(file, nextIndex)
	}
	//fmt.Println(file)
	startIndex := 0
	for i := 0; i < len(file)-1; i++ {
		if file[i] == 0 {
			startIndex = i
			break
		}
	}
	fmt.Println("Start index", startIndex)

	sum := 0
	sum += file[startIndex+1000]
	sum += file[startIndex+2000]
	sum += file[startIndex+3000]

	/*
	for i := 0; i < 3000; i++ {
		startIndex++
		if startIndex > len(file)-1 {
			startIndex = 0;
		}
		if i % 1000 == 0 {
			fmt.Printf("%d = %d\n", i , file[startIndex])
			sum += file[startIndex]
		}


	}
	*/

	return sum
}

func task2(lines []string) {

}

func parseLines(lines []string) []int {
	var out []int
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		out = append(out, n)
	}
	return out
}

func moveValue(file []int, index int) []int {
	moves := int(math.Abs(float64(file[index])))
	delta := 1
	if file[index] < 0 {
		delta = -1
	}

	for moves != 0 {
		moves--

		targetIndex := index + delta

		if targetIndex <= 0 && delta < 0 {
			if targetIndex == 0 {
				file[index], file[targetIndex] = file[targetIndex], file[index]
			}
			tmp := file[0]
			file = file[1:]
			file = append(file, tmp)
			index = len(file)-1
			//moves++
			continue
		} else if targetIndex >= len(file)-1 && delta > 0 {
			if targetIndex == len(file)-1 {
				file[index], file[targetIndex] = file[targetIndex], file[index]
			}
			tmp := file[len(file)-1]
			file = file[0:len(file)-1]
			file = append([]int{tmp}, file...)
			index = 0
			continue
		}

		/*if targetIndex < 0 {
			targetIndex = len(file)-1
		} else if targetIndex > len(file)-1 {
			targetIndex = 0
		}*/
		file[index], file[targetIndex] = file[targetIndex], file[index]
		index += delta
		
		
		
	}
	return file
}

func check(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}