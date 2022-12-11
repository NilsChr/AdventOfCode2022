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
	lines := utils.GetInput("./day11/input-test.txt")
	fmt.Println("Task 1:", task1(parseMonkeys(lines)))
	fmt.Println("Task 2:", task2(parseMonkeys(lines)))
}

type Monkey struct {
	id            int
	inspects      uint64
	items         []uint64
	operator      string
	operatorValue uint64
	operatorSame  bool
	testValue     uint64
	testTrue      int
	testFalse     int
}

func (m *Monkey) inspect(item uint64) uint64 {
	var out uint64 = 0
	if m.operatorSame {
		m.operatorValue = item
	}
	switch m.operator {
	case "*":
		out = m.operatorValue * item
	case "+":
		out = m.operatorValue + item
	}
	//fmt.Println(out)
	m.inspects++
	return out
}

func parseMonkeys(lines []string) []Monkey {
	var monkeys []Monkey

	for i := 0; i < len(lines); i += 7 {
		//fmt.Println(lines[i])
		parts := strings.Split(lines[i], " ")
		monkey := new(Monkey)
		monkey.id, _ = strconv.Atoi(strings.Replace(parts[1], ":", "", 1))

		// Parse Items
		itemLine := lines[i+1]
		itemLine = strings.Replace(itemLine, "Starting items: ", "", 1)
		items := strings.Split(itemLine, ",")
		for j := 0; j < len(items); j++ {
			item, _ := strconv.Atoi(strings.TrimSpace(items[j]))
			monkey.items = append(monkey.items, uint64(item))
		}

		// Parse Operator
		operatorLine := lines[i+2]
		operatorLine = strings.Replace(operatorLine, "Operation: new = old", "", 1)
		operatorLine = strings.TrimSpace(operatorLine) //:= strings.Split(operatorLine, " ")
		operatorParts := strings.Split(operatorLine, " ")
		monkey.operator = operatorParts[0]
		parsed, err := strconv.Atoi(strings.TrimSpace(operatorParts[1]))
		if err != nil {
			monkey.operatorSame = true
		} else {
			monkey.operatorValue = uint64(parsed)
		}

		// Parse test
		testLine := lines[i+3]
		testParts := strings.Split(testLine, " ")
		p, _ := strconv.Atoi(testParts[len(testParts)-1])
		monkey.testValue = uint64(p)

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

func task1(monkeys []Monkey) uint64 {
	//fmt.Println(monkeys)

	for i := 0; i < 20; i++ {
		for m := 0; m < len(monkeys); m++ {
			monkey := monkeys[m]
			if len(monkeys[m].items) == 0 {
				continue
			}
			for {
				if len(monkeys[m].items) == 0 {
					break
				}
				//fmt.Println("")
				//fmt.Println("")

				item := monkeys[m].items[:1]
				monkeys[m].items = monkeys[m].items[1:]
				//fmt.Printf("Monkey %d inspects %d\n", m, item[0])
				//fmt.Println("After pick:", monkeys[m].items)

				itemAfter := monkeys[m].inspect(item[0])
				//fmt.Printf("Level is multiplied by %d to %d\n", monkey.operatorValue, itemAfter)

				rounded := uint64(math.Floor(float64(itemAfter) / 3))
				//fmt.Printf("Divided by 3 to %d\n", rounded)
				testPassed := rounded%uint64(monkeys[m].testValue) == 0
				//fmt.Printf("Test passed: %t \n", testPassed)
				if testPassed {
					//fmt.Printf("Monkey %d before %v\n", monkey.testTrue, monkeys[monkey.testTrue].items)
					monkeys[monkey.testTrue].items = append(monkeys[monkey.testTrue].items, rounded)
					//fmt.Printf("Item with worry level %d is passed to monkey %d\n", rounded, monkey.testTrue)
					//fmt.Printf("Monkey %d after %v\n", monkey.testTrue, monkeys[monkey.testTrue].items)

				} else {
					//fmt.Printf("Monkey %d before %v\n", monkey.testFalse, monkeys[monkey.testFalse].items)
					monkeys[monkey.testFalse].items = append(monkeys[monkey.testFalse].items, rounded)
					//fmt.Printf("Item with worry level %d is passed to monkey %d\n", rounded, monkey.testFalse)
					//fmt.Printf("Monkey %d after %v\n", monkey.testFalse, monkeys[monkey.testFalse].items)

				}
			}

		}
	}
	//fmt.Println(monkeys)

	sort.Slice(monkeys, func(i,j int) bool {
		return monkeys[i].inspects > monkeys[j].inspects
	})

	return monkeys[0].inspects * monkeys[1].inspects
}

// Too high 12911266280

// Test 
// 2637590098 - my
// 2713310158 - correct
func task2(monkeys []Monkey) uint64 {
	//fmt.Println(monkeys)

	for i := 1; i <= 10000; i++ {
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
				testPassed := itemAfter % uint64(monkeys[m].testValue) == 0
				if testPassed {
					monkeys[monkeys[m].testTrue].items = append(monkeys[monkeys[m].testTrue].items, itemAfter)
				} else {
					monkeys[monkeys[m].testFalse].items = append(monkeys[monkeys[m].testFalse].items, itemAfter)
				}
			}

		}

		if i == 1000 {
			printWorries(monkeys,i)
		}
	}
	sort.Slice(monkeys, func(i,j int) bool {
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