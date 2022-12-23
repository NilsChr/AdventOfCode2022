package day23

import (
	"advent-of-code-2022/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	lines := utils.GetInput("./input-test2.txt")
	elfes, dim := parseInput(lines)

	assert.Equal(t,len(elfes), 1, "It should be exactly one elf")
	assert.Equal(t,dim.X, 5, "Dim X should be 5")
	assert.Equal(t,dim.Y, 5, "Dim Y should be 5")
	assert.Equal(t,elfes[0].pos.X, 2, "Elf pos X should be 2")
	assert.Equal(t,elfes[0].pos.Y, 2, "Elf pos Y should be 2")

	elfes[0].considerMove(elfes, 0)
	assert.Equal(t,elfes[0].proposedPos.X, 2, "Elf proppos X should be 2")
	assert.Equal(t,elfes[0].proposedPos.Y, 1, "Elf proppos Y should be 1")
	uniqueMoves := make(map[utils.Vec2]int)
	uniqueMoves[elfes[0].proposedPos]++
	assert.Equal(t,len(uniqueMoves), 1, "Should find one proposed position")
	assert.Equal(t,uniqueMoves[elfes[0].proposedPos], 1, "Should find count == 1 on propsed pos")

	if uniqueMoves[elfes[0].proposedPos] == 1 {
		fmt.Println("SHOUD MOVE")
		fmt.Println("FROM",  elfes[0].pos)
		fmt.Println("TO",  elfes[0].proposedPos)
		elfes[0].pos = *elfes[0].proposedPos.Copy()
		//elfes[0].pos.X = elfes[0].proposedPos.X
		//elfes[0].pos.Y = elfes[0].proposedPos.Y

		fmt.Println(elfes[0].pos)
	}
	assert.Equal(t,elfes[0].pos.X, 2, "Elf pos X should be 2")
	assert.Equal(t,elfes[0].pos.Y, 1, "Elf pos Y should be 1")
}