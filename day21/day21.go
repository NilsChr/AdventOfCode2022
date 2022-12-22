package day21

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
	"strings"
)

func Day21() {
	lines := utils.GetInput("./day21/input.txt")
	fmt.Println("Task1: ", task1(lines))
	fmt.Println("Task2: ", task2(lines))

}

func task1(lines []string) int {
	monkeys := parseMonkeys(lines)
	return solveMonkey("root", monkeys)
}

func task2(lines []string) int {
	monkeys := parseMonkeys(lines)
	root := monkeys["root"]

	rootA := 1
	rootB := 2

	/*const MaxUint = ^uint(0)
	const MinUint = 0
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1
	*/

	//rootA = solveMonkey(root.nextA, monkeys)
	//rootB = solveMonkey(root.nextB, monkeys)
	//	fmt.Println(rootA, rootB)
	//return solveMonkey("root", monkeys)
	start := 0//math.MinInt64//MinInt
	//end   := 0//math.MaxInt64//MinInt
	mid := -1
	i := start
	for rootA != rootB {
		i++
		//mid = int((end + start) / 2)

		
		me := monkeys["humn"]
		//me.solved = true
		me.value = i
		monkeys["humn"] = me
		rootA = solveMonkey(root.nextA, monkeys)
		rootB = solveMonkey(root.nextB, monkeys)
 	
	}

	fmt.Printf("v:%d, a:%d b:%d\n", mid, rootA, rootB)

	return mid //solveMonkey("root", monkeys)
}

type monkey struct {
	id       string
	//solved   bool
	value    int
	nextA    string
	nextB    string
	operator string
}

func solveMonkey(start string, monkeys map[string]*monkey) int {
	monkey := monkeys[start]
	//if !monkey.solved {
	switch monkey.operator {
	case "+":
		monkey.value = solveMonkey(monkey.nextA, monkeys) + solveMonkey(monkey.nextB, monkeys)
	case "-":
		monkey.value = solveMonkey(monkey.nextA, monkeys) - solveMonkey(monkey.nextB, monkeys)
	case "*":
		monkey.value = solveMonkey(monkey.nextA, monkeys) * solveMonkey(monkey.nextB, monkeys)
	case "/":
		monkey.value = solveMonkey(monkey.nextA, monkeys) / solveMonkey(monkey.nextB, monkeys)
	}
	//}
	//monkey.solved = true
	monkeys[monkey.id] = monkey
	return monkey.value
}

func parseMonkeys(lines []string) map[string]*monkey {
	monkeys := make(map[string]*monkey)

	for _, line := range lines {
		parts := strings.Split(line, " ")
		id := strings.Replace(parts[0], ":", "", 1)
		monkey := new(monkey)
		monkey.id = id

		if len(parts) == 2 {
			monkey.value, _ = strconv.Atoi(parts[1])
		} else {
			monkey.nextA = parts[1]
			monkey.operator = parts[2]
			monkey.nextB = parts[3]
		}
		monkeys[id] = monkey
	}

	return monkeys
}


