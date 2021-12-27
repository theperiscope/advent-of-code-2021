package main

import (
	"AOC/pkg/utils"
	"fmt"
	"os"
	"strings"
	"unsafe"
)

type grid [][]int

type point struct {
	X int
	Y int
}

func (g *grid) Print() {
	for i := 0; i < len(*g); i++ {
		for j := 0; j < len((*g)[i]); j++ {
			if (*g)[i][j] == 0 {
				fmt.Printf("\x1b[104m%d\x1b[0m", (*g)[i][j])
			} else {
				fmt.Printf("\x1b[90m%d\x1b[0m", (*g)[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *grid) Neighbors(i, j int) []point {
	neighbors := []point{}
	if i-1 >= 0 {
		neighbors = append(neighbors, point{Y: i - 1, X: j})
	}
	if i+1 <= len(*g)-1 {
		neighbors = append(neighbors, point{Y: i + 1, X: j})
	}
	if j-1 >= 0 {
		neighbors = append(neighbors, point{Y: i, X: j - 1})
	}
	if j+1 <= len((*g)[i])-1 {
		neighbors = append(neighbors, point{Y: i, X: j + 1})
	}
	if i-1 >= 0 && j-1 >= 0 {
		neighbors = append(neighbors, point{Y: i - 1, X: j - 1})
	}
	if i+1 <= len(*g)-1 && j-1 >= 0 {
		neighbors = append(neighbors, point{Y: i + 1, X: j - 1})
	}
	if i-1 >= 0 && j+1 <= len((*g)[i])-1 {
		neighbors = append(neighbors, point{Y: i - 1, X: j + 1})
	}
	if i+1 <= len(*g)-1 && j+1 <= len((*g)[i])-1 {
		neighbors = append(neighbors, point{Y: i + 1, X: j + 1})
	}
	return neighbors
}

func (g *grid) Step() (count int) {
	flashed := make([][]bool, len(*g))

	for i := 0; i < len(*g); i++ {
		flashed[i] = make([]bool, len((*g)[i]))
	}

	flashPoints := []point{}
	for i := 0; i < len(*g); i++ {
		for j := 0; j < len((*g)[i]); j++ {
			(*g)[i][j]++

			if (*g)[i][j] > 9 && !flashed[i][j] {
				flashPoints = append(flashPoints, point{j, i})
				flashed[i][j] = true
				(*g)[i][j] = 0
			}
		}
	}

	for len(flashPoints) > 0 {
		p := flashPoints[0]
		flashPoints = flashPoints[1:]

		neighbors := g.Neighbors(p.Y, p.X)
		for _, n := range neighbors {
			(*g)[n.Y][n.X]++
			if (*g)[n.Y][n.X] > 9 && !flashed[n.Y][n.X] {
				flashPoints = append(flashPoints, n)
				flashed[n.Y][n.X] = true
				(*g)[n.Y][n.X] = 0
			}
		}
	}

	count = 0
	for i := 0; i < len(*g); i++ {
		for j := 0; j < len((*g)[i]); j++ {
			if flashed[i][j] {
				count++
				(*g)[i][j] = 0
			}
		}
	}

	return count
}

func part1(g *grid) {
	count := 0
	for i := 0; i < 100; i++ {
		count += g.Step()
	}

	fmt.Println("Part 1 Answer:", count)
}

func part2(g *grid) {
	i := 1
	flashesCount := g.Step()
	for flashesCount != 100 {
		flashesCount = g.Step()
		i++
	}

	fmt.Println("Part 2 Answer:", i)
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Printf("Usage: %s <inputfile>\n", utils.GetProgramName())
		return
	}

	lines, _ := utils.ReadInput(argsWithoutProg[0])
	g1 := [][]int{}
	g2 := [][]int{}
	for _, line := range lines {
		g1 = append(g1, utils.StringToInt(strings.Split(line, "")))
		g2 = append(g2, utils.StringToInt(strings.Split(line, "")))
	}

	// per https://stackoverflow.com/questions/29031353/conversion-of-a-slice-of-string-into-a-slice-of-custom-type
	part1((*grid)(unsafe.Pointer(&g1)))
	part2((*grid)(unsafe.Pointer(&g2)))
}
