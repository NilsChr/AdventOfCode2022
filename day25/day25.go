package day25

import (
	"advent-of-code-2022/utils"
	"fmt"
	"math"
)

func Day25() {
	lines := utils.GetInput("./day25/input.txt")
	fmt.Println("Task1: ", task1(lines))
}

func task1(lines []string) string {

	sum := 0
	for _, line := range lines {
		sum += decodeNAFU(line)
	}

	return 	encodeSNAFU(sum)
}

func decodeNAFU(in string) int {
	sum := 0
	for i := range in {
		sum += addSNAFU(string(in[i]), len(in)-i-1)
	}
	return sum
}

func addSNAFU(in string, space int) int {
	mult := int(math.Pow(5, float64(space)))
	if space == 0 {
		mult = 1
	}
	switch in {
	case "=": {
			return -2 * mult
		}
	case "-":{
			return -1 * mult
		}
	case "0":{ 
			return 0 * mult
		}
	case "1":{
			return 1 * mult
		}
	case "2":{
			return 2 * mult
		}
	}
	return 0
}


func toSnafuDigit(in int) string {
	switch in {
		case -2 : {
			return "="
			}
		case -1:{
				return "-"
			}
		case 0:{ 
				return "0"
			}
		case 1:{
				return "1"
			}
		case 2:{
				return "2"
			}
		}
		return ""
}
  
func encodeSNAFU(in int) string {
	var digits []int
	for in > 0 {
		digits = append(digits, (in + 2) % 5)
		in = int(math.Floor(float64(in + 2)) / 5)
	}

	snafu := ""
	for i := len(digits)-1; i >= 0; i-- {
		snafu += toSnafuDigit(digits[i]-2)
	}
	return snafu

}