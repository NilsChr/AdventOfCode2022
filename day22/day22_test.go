package day22

import (
//	"advent-of-code-2022/utils"
	"fmt"
	"testing"
)


func Test1(t *testing.T) {
	//lines := utils.GetInput("./input-test4.txt")
	//grid, commands := parseInput(lines)
    //pos := utils.NewVec2(3,2)
	//side := getCurrentSide(2, pos)
	//fmt.Println(side)

	x, y := moveCoordinateOnCubeTexture(3, 0,2,2, 0)
	fmt.Println(x,y)

}
/*

    0       1		2
	0	1	2	3	4	5

0   		

1   

2

3
*/