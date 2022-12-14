package day14

import (
	"advent-of-code-2022/utils"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	fmt.Println("Hello")

	walls := make(map[utils.Vec2]bool)
	p1 := utils.NewVec2(1,1)
	p2 := utils.NewVec2(1,1)

	walls[*p1] = true

	actual := walls[*p1]
	assert.Equal(t, actual, true, "The position should be marked true")

	actual2 := walls[*p2]
	assert.Equal(t, actual2, true, "The position should be marked true")
}