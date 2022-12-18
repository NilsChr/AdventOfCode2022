package day12

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func printGrid(grid [][]Node) {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	row := ""
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			row += string(rune(grid[y][x].value))
		}
		fmt.Println(row)
		row = ""
	}

}

func printGrid2(grid [][]Node, focus Node, look Node, visited []Node, queue []Node, neighbours []Vec2) {
	//cmd := exec.Command("cmd", "/c", "cls")
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	/*
		fmt.Println("Focus:", focus, grid[focus.y][focus.x].value, string(rune(grid[focus.y][focus.x].value)))
		fmt.Println("Looking at:", look, grid[look.y][look.x].value, string(rune(grid[look.y][look.x].value)))

		fmt.Println("Visited", visited)
		fmt.Println("Queue", queue)
		fmt.Println("Neighbours", neighbours)
		fmt.Println("")
	*/
	row := ""
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			val := ""
			if focus.x == x && focus.y == y {
				//row += fmt.Sprint("\033[35m" + "*" + "\033[0m")
				val = fmt.Sprint("\033[35m" + string(rune(grid[y][x].value)) + "\033[0m")
			} else if look.x == x && look.y == y {
				//row += fmt.Sprint("\033[32m" + "¤" + "\033[0m")
				val = fmt.Sprint("\033[32m" + string(rune(grid[y][x].value)) + "\033[0m")
			} else {
				val = string(rune(grid[y][x].value))
				if Contains(visited, grid[y][x]) {
					val = fmt.Sprint("\033[22m" + "." + "\033[0m")
				}
			}

			row += val

		}
		fmt.Println(row)
		row = ""
	}
	fmt.Println("QUEUE:", len(queue))
	fmt.Println("VISITED:", len(visited))
	//time.Sleep(1 * time.Second)
	//fmt.Scanln()
	wait()
}

func printGrid3(grid [][]Node, focus Node, visited []Node, queue []Node) {
	//cmd := exec.Command("cmd", "/c", "cls")
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	/*
		fmt.Println("Focus:", focus, grid[focus.y][focus.x].value, string(rune(grid[focus.y][focus.x].value)))
		fmt.Println("Looking at:", look, grid[look.y][look.x].value, string(rune(grid[look.y][look.x].value)))

		fmt.Println("Visited", visited)
		fmt.Println("Queue", queue)
		fmt.Println("Neighbours", neighbours)
		fmt.Println("")
	*/
	row := ""
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			val := ""
			if focus.x == x && focus.y == y {
				//row += fmt.Sprint("\033[35m" + "*" + "\033[0m")
				val = fmt.Sprint("\033[35m" + string(rune(grid[y][x].value)) + "\033[0m")
			} else {
				val = string(rune(grid[y][x].value))
				if Contains(visited, grid[y][x]) {
					val = fmt.Sprint("\033[22m" + "." + "\033[0m")
				}
			}

			row += val

		}
		fmt.Println(row)
		row = ""
	}

	fmt.Println("QUEUE:", len(queue))
	fmt.Println("VISITED:", len(visited))

	//time.Sleep(1 * time.Second)
	//fmt.Scanln()
	wait()
}

func wait() {
	//fmt.Scanln()
	time.Sleep(1 * time.Nanosecond)
}
