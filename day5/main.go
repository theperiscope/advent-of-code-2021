package main

import (
	"AOC/pkg/utils"
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Printf("Usage: %s <inputfile>\n", utils.GetProgramName())
		return
	}

	lines, _ := utils.ReadInput(argsWithoutProg[0])

	part1 := utils.Filter(lines, func(s string) bool {
		var x1, y1, x2, y2 int
		fmt.Sscanf(s, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		return x1 == x2 || y1 == y2
	})

	part2 := lines

	work(part1)
	work(part2)

}

func work(part2 []string) {
	var grid [1000][1000]int
	for _, line := range part2 {

		var x1, y1, x2, y2 int
		fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)

		t := ""
		if x1 == x2 {
			t = "h"
		} else if y1 == y2 {
			t = "v"
		} else {
			t = "45"
		}

		switch t {
		case "45":
			dx := 0
			dy := 0
			N := 0

			if x2-x1 > 0 {
				dx = 1
				N = x2 - x1
			} else {
				dx = -1
				N = x1 - x2
			}
			if y2-y1 > 0 {
				dy = 1
			} else {
				dy = -1
			}

			//fmt.Println(x1, y1, dx, dy)
			for n := 0; n <= N; n++ {
				grid[x1+(n*dx)][y1+(n*dy)]++
			}
		case "h":
			N := 0
			startY := y1
			if y1 > y2 {
				N = y1 - y2
				startY = y2
			} else if y1 < y2 {
				N = y2 - y1
				startY = y1
			}
			for n := 0; n <= N; n++ {
				grid[x1][startY+n]++
			}
		case "v":
			N := 0
			startX := x1
			if x1 > x2 {
				N = x1 - x2
				startX = x2
			} else if x1 < x2 {
				N = x2 - x1
				startX = x1
			}
			for n := 0; n <= N; n++ {
				grid[startX+n][y1]++
			}
		}
	}

	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[j][i] >= 2 {
				count++
			}
		}
	}

	fmt.Println(count)
}
