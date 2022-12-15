package day13

import (
	"advent-of-code-2022/utils"
	"fmt"
	"strconv"
)


type tree struct {
	value int
	children []*tree
	parent *tree
}

func Day13_2() {
	lines := utils.GetInput("./day13/input-test.txt")
	a := parseTree(lines[3])
	b := parseTree(lines[4])

	fmt.Println(areOrdered(a,b))
}



func parseTree(line string) tree {
	root := tree{-1, []*tree{}, nil}
	temp := &root

	var currentNumber string
	for _,r := range line {
		switch r {
		case '[':
			newTree := tree{-1, []*tree{}, temp}
			temp.children=append(temp.children, &newTree)
			temp = &newTree
		case ']':
			if len(currentNumber) >= 0 {
				number, _ := strconv.Atoi(currentNumber)
				temp.value = number
				currentNumber =""
			}
			temp = temp.parent
		case ',':
			if len(currentNumber) >= 0 {
				number, _ := strconv.Atoi(currentNumber)
				temp.value = number
				currentNumber =""
			}
			temp = temp.parent
			newTree := tree{-1, []*tree{}, temp}
			temp.children=append(temp.children, &newTree)
			temp = &newTree
		default : 
			currentNumber += string(r)
		}
	}
	return root
}

func areOrdered(first, second tree)int{
	switch{
	case len(first.children) == 0 && len(second.children) == 0:
		if first.value > second.value{
			return -1
		} else if first.value == second.value {
			return 0
		}
		return 1
		
	case first.value >= 0:
		return areOrdered(tree{-1, []*tree{&first}, nil}, second)

	case second.value >= 0:
		return areOrdered(first, tree{-1, []*tree{&second}, nil})
	default:
		var i int
		for i=0; i<len(first.children) && i<len(second.children); i++{
			ordered := areOrdered(*first.children[i], *second.children[i])
			if ordered != 0{
				return ordered
			}
		}
		if i < len(first.children){
			return -1
		}else if i < len(second.children){
			return 1
		}
	}
	return 0
}