package day11

import (
	"advent-of-code-2022/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day11() {
	lines := utils.GetInput("./day11/input.txt")
	fmt.Println("Task 1:", task(parseMonkeys(lines),20,1))
	fmt.Println("Task 2:", task(parseMonkeys(lines),10000,2))
}

type Monkey struct {
	id            int
	inspects      int
	items         []int
	operator      string
	operatorValue int
	operatorSame  bool
	testValue     int
	testTrue      int
	testFalse     int
}

func (m *Monkey) inspect(item int) int {
	var out int = 0
	if m.operatorSame {
		m.operatorValue = item
	}
	switch m.operator {
	case "*":
		out = m.operatorValue * item
	case "+":
		out = m.operatorValue + item
	}
	m.inspects++
	return out
}

func parseMonkeys(lines []string) []Monkey {
	var monkeys []Monkey

	for i := 0; i < len(lines); i += 7 {
		parts := strings.Split(lines[i], " ")
		monkey := new(Monkey)
		monkey.id, _ = strconv.Atoi(strings.Replace(parts[1], ":", "", 1))

		// Parse Items
		itemLine := lines[i+1]
		itemLine = strings.Replace(itemLine, "Starting items: ", "", 1)
		items := strings.Split(itemLine, ",")
		for j := 0; j < len(items); j++ {
			item, _ := strconv.Atoi(strings.TrimSpace(items[j]))
			monkey.items = append(monkey.items, int(item))
		}

		// Parse Operator
		operatorLine := lines[i+2]
		operatorLine = strings.Replace(operatorLine, "Operation: new = old", "", 1)
		operatorLine = strings.TrimSpace(operatorLine)
		operatorParts := strings.Split(operatorLine, " ")
		monkey.operator = operatorParts[0]
		parsed, err := strconv.Atoi(strings.TrimSpace(operatorParts[1]))
		if err != nil {
			monkey.operatorSame = true
		} else {
			monkey.operatorValue = int(parsed)
		}

		// Parse test
		testLine := lines[i+3]
		testParts := strings.Split(testLine, " ")
		p, _ := strconv.Atoi(testParts[len(testParts)-1])
		monkey.testValue = int(p)

		// Parse test true
		testLineT := lines[i+4]
		testPartsT := strings.Split(testLineT, " ")
		monkey.testTrue, _ = strconv.Atoi(testPartsT[len(testPartsT)-1])

		// Parse test true
		testLineF := lines[i+5]
		testPartsF := strings.Split(testLineF, " ")
		monkey.testFalse, _ = strconv.Atoi(testPartsF[len(testPartsF)-1])

		monkeys = append(monkeys, *monkey)
	}

	return monkeys
}

func task(monkeys []Monkey, runs int,part int) int {
	supermod := 1

	for _, m := range monkeys {
		supermod *= m.testValue
	}

	for i := 1; i <= runs; i++ {
		for m := 0; m < len(monkeys); m++ {
			if len(monkeys[m].items) == 0 {
				continue
			}
			for {
				if len(monkeys[m].items) == 0 {
					break
				}
				item := monkeys[m].items[:1]
				monkeys[m].items = monkeys[m].items[1:]
				itemAfter := monkeys[m].inspect(item[0])
				if part == 1 {
					itemAfter = int(math.Floor(float64(itemAfter) / 3))
				} else {
					itemAfter = itemAfter % supermod
				}
				testPassed := itemAfter%int(monkeys[m].testValue) == 0
				if testPassed {
					monkeys[monkeys[m].testTrue].items = append(monkeys[monkeys[m].testTrue].items, itemAfter)
				} else {
					monkeys[monkeys[m].testFalse].items = append(monkeys[monkeys[m].testFalse].items, itemAfter)
				}
			}

		}

	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspects > monkeys[j].inspects
	})
	return monkeys[0].inspects * monkeys[1].inspects
}

func printWorries(monkeys []Monkey, round int) {
	fmt.Printf("After round %d\n", round)
	for _, monkey := range monkeys {
		fmt.Printf("Monkey %d inspected items %d times.\n", monkey.id, monkey.inspects)
	}
}