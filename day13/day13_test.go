package day13

import (
	"fmt"
	"sort"
	"testing"
)

func Test1(t *testing.T) {
	t.Skip()
	//[[],[8,10,6],[3]]
	//[[[3],4],[10,[],[]]]

	//input := "[[8,7,6]]"
	//inputB := "[8,7,6]"

	

	input := "[]"
	inputB := "[1]"

	partAEntropy := 0

	var unorderedA []int
	var unorderedB []int

	for input != "" {
		partAEntropy++
		//fmt.Println("Consuming: ", input)
		output, consumedSlice := consumeLine(input)
		ordered := sort.IntsAreSorted(consumedSlice)
		if(!ordered) {
			//fmt.Println("Part A not ordered", consumedSlice)
			//break
			unorderedA = append(unorderedA, partAEntropy)
		}
		//fmt.Println(output, consumedSlice, "Ordered,", ordered)
		input = output
		
	}


	partBEntropy := 0

	for inputB != "" {
		partBEntropy++

		//fmt.Println("Consuming: ", inputB)
		output, consumedSlice := consumeLine(inputB)
		ordered := sort.IntsAreSorted(consumedSlice)
		if(!ordered) {
			//fmt.Println("Part B not ordered", consumedSlice)
			unorderedB = append(unorderedB, partAEntropy)

		}
		//fmt.Println(output, consumedSlice, "Ordered,", ordered)
		inputB = output
	}

	fmt.Println("part a unordered:", unorderedA)
	fmt.Println("part b unordered:", unorderedB)


	// b unordered len > a unordered len : inputs are not in the right order
}

func Test2(t *testing.T) {
	t.Skip()

	a := "[1,[2,[3,[4,[5,6,7]]]],8,9]"
	b := "[1,[2,[3,[4,[5,6,0]]]],8,9]"

	fmt.Println("Reduced A:", reduceTo1D(a))
	fmt.Println("Reduced B:", reduceTo1D(b))
}


func Test3(t *testing.T) {
	t.Skip()
	//a := "[[[]][]]"
	//b := "[[][][[[]]]]"
	a := "[[[],[[],[2,3,9]],5,5,5],[[1],5,[[],4],[5,[3,9,3],[6,4,1]],3],[3,[4,1]],[[],[4,[8],[7]],[[10,8,2],5]],[[[1,5,0,9,6],9,7,[]],[0,[],6],[],[[3,5,8],6]]]"
    b := "[[[8,[],8,[7],3]]]"

	fmt.Println("Reduced A:", reduceTo1D(a))
	fmt.Println("Reduced B:", reduceTo1D(b))
}

func Test4(t *testing.T) {
	t.Skip()
	/*
	var test []int
	res := collapseSlice(test)
	fmt.Println(res)
	*/

	sliceString := stringToSlice("[]")
	fmt.Println(sliceString)
}

func Test5(t *testing.T) {
	var a []int = []int{1,2,3}
	var b []int = []int{1,4,3}

	order := checkSliceOrder(a,b)
	fmt.Println(order)
}