package main

import (
	"AOC/pkg/utils"
	"fmt"
	"strings"
)

type grid [][]string

func (g *grid) Print() {
	for i := 0; i < len(*g); i++ {
		for j := 0; j < len((*g)[i]); j++ {
			if (*g)[i][j] == ">" || (*g)[i][j] == "v" {
				fmt.Printf("\x1b[104m%s\x1b[0m", (*g)[i][j])
			} else {
				fmt.Printf("\x1b[90m%s\x1b[0m", (*g)[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *grid) Step() int {
	count, w, h := 0, len((*g)[0]), len((*g))
	whichCanMove := [][]int{}
	for y := h - 1; y >= 0; y-- {
		for x := w - 1; x >= 0; x-- {
			if (*g)[y][x] == ">" {
				nX := (x + 1) % w
				if (*g)[y][nX] == "." {
					whichCanMove = append(whichCanMove, []int{y, x})
				}
			}
		}
	}
	count += len(whichCanMove)
	for i := 0; i < len(whichCanMove); i++ {
		c := whichCanMove[i]
		switch (*g)[c[0]][c[1]] {
		case ">":
			nX := (c[1] + 1) % w
			(*g)[c[0]][c[1]] = "."
			(*g)[c[0]][nX] = ">"
		case "v":
			nY := (c[0] + 1) % h
			(*g)[c[0]][c[1]] = "."
			(*g)[nY][c[1]] = "v"
		}
	}
	whichCanMove = [][]int{}
	for x := w - 1; x >= 0; x-- {
		for y := h - 1; y >= 0; y-- {
			if (*g)[y][x] == "v" {
				nY := (y + 1) % h
				if (*g)[nY][x] == "." {
					whichCanMove = append(whichCanMove, []int{y, x})
				}
			}
		}
	}
	count += len(whichCanMove)
	for i := 0; i < len(whichCanMove); i++ {
		c := whichCanMove[i]
		switch (*g)[c[0]][c[1]] {
		case ">":
			nX := (c[1] + 1) % w
			(*g)[c[0]][c[1]] = "."
			(*g)[c[0]][nX] = ">"
		case "v":
			nY := (c[0] + 1) % h
			(*g)[c[0]][c[1]] = "."
			(*g)[nY][c[1]] = "v"
		}
	}
	return count
}

func part1(g1 grid) {
	//g1.Print()
	i := 1
	for {
		count := g1.Step()
		//g1.Print()
		fmt.Println("Step", i, "moved", count, "sea cucumbers.")
		if count == 0 {
			fmt.Println("Done.")
			break
		}
		i++
	}
}

func main() {
	utils.AssertArgs()
	lines := utils.AssertInput()
	g1 := [][]string{}
	for _, line := range lines {
		g1 = append(g1, strings.Split(line, ""))
	}

	part1(g1)
}
