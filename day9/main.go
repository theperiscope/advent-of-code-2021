package main

import (
	"AOC/pkg/utils"
	"fmt"
	"os"
	"sort"
	"strings"
)

type point struct {
	X int
	Y int
}

func part1(numbers [][]int) (lowPoints []point) {
	lowPoints, totalRisk := []point{}, 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers[i]); j++ {
			n, up, down, left, right := numbers[i][j], 9, 9, 9, 9
			if i-1 >= 0 {
				up = numbers[i-1][j]
			}
			if i+1 <= len(numbers)-1 {
				down = numbers[i+1][j]
			}
			if j-1 >= 0 {
				left = numbers[i][j-1]
			}
			if j+1 <= len(numbers[i])-1 {
				right = numbers[i][j+1]
			}

			// low point is defined by bigger numbers on all 4 sides
			if n < up && n < down && n < left && n < right {
				lowPoints = append(lowPoints, point{X: j, Y: i})
				totalRisk += 1 + n
			}
		}
	}

	fmt.Println("Part 1 Answer:", totalRisk)
	return
}

func part2(numbers [][]int, lowPoints []point) {
	basinSizes := []int{}
	for _, lowPoint := range lowPoints {
		basinSizes = append(basinSizes, basinSize(lowPoint, numbers))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	fmt.Println("Part 2 Answer:", basinSizes[0]*basinSizes[1]*basinSizes[2])
}

func basinSize(lowPoint point, input [][]int) int {
	numbers := utils.CloneSliceIntInt(input) // clone because we need to modify it during processing below

	// start from lowPoint
	// basin is surrounded by ever-increasing neighbors or 9s or until we run out of bounds
	queue := []point{lowPoint}
	size := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if numbers[current.Y][current.X] >= 0 {
			size++                             // count only if not already counted
			numbers[current.Y][current.X] = -1 // -1 represents visited state
		}

		if current.Y-1 >= 0 && numbers[current.Y-1][current.X] >= 0 && numbers[current.Y-1][current.X] != 9 && numbers[current.Y-1][current.X] > numbers[current.Y][current.X] { // top neighbor
			queue = append(queue, point{X: current.X, Y: current.Y - 1})
		}
		if current.Y+1 <= len(numbers)-1 && numbers[current.Y+1][current.X] >= 0 && numbers[current.Y+1][current.X] != 9 && numbers[current.Y+1][current.X] > numbers[current.Y][current.X] { // bottom neighbor
			queue = append(queue, point{X: current.X, Y: current.Y + 1})
		}
		if current.X-1 >= 0 && numbers[current.Y][current.X-1] >= 0 && numbers[current.Y][current.X-1] != 9 && numbers[current.Y][current.X-1] > numbers[current.Y][current.X] { // left neighbor
			queue = append(queue, point{X: current.X - 1, Y: current.Y})
		}
		if current.X+1 <= len(numbers[current.Y])-1 && numbers[current.Y][current.X+1] >= 0 && numbers[current.Y][current.X+1] != 9 && numbers[current.Y][current.X+1] > numbers[current.Y][current.X] { // right neighbor
			queue = append(queue, point{X: current.X + 1, Y: current.Y})
		}
	}
	return size
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Printf("Usage: %s <inputfile>\n", utils.GetProgramName())
		return
	}

	lines, _ := utils.ReadInput(argsWithoutProg[0])
	numbers := [][]int{}
	for _, line := range lines {
		numbers = append(numbers, utils.StringToInt(strings.Split(line, "")))
	}

	lowPoints := part1(numbers)
	part2(numbers, lowPoints)
}
