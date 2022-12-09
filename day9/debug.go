package day9

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func debug(rope []Vec2, positions map[string]bool) {
	cmd := exec.Command("clear") 
	cmd.Stdout = os.Stdout
    cmd.Run()
	fmt.Println(rope)
	ox := 11
	oy := 16
	for y := 0; y < 22; y++ {
		for x := 0; x < 26; x++ {
			key := fmt.Sprintf("%d,%d", x-ox,y-oy)
			rx := rope[0].x
			ry := rope[0].y

			if positions[key] {
				fmt.Print("#")
			} else if rx == x -ox && ry == y-oy {
				fmt.Print("H")
			} else if rx == x -ox && ry == y-oy {
				fmt.Print("T")
			} else {
				fmt.Print(".")
			}
			
		}
		fmt.Println("")
	}
	time.Sleep(50 * time.Millisecond)
}