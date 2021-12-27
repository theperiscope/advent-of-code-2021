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

func (p point) String() string {
	return fmt.Sprintf("(X=%d,Y=%d)", p.X, p.Y)
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
	return neighbors
}

func part1(g *grid) {
	graph := utils.Graph{}
	graph.Directed = true

	for i := 0; i < len(*g); i++ {
		for j := 0; j < len((*g)[i]); j++ {
			node := fmt.Sprintf("N%03X%03X", i, j)
			neighbors := g.Neighbors(i, j)
			for _, n := range neighbors {
				neighborNode := fmt.Sprintf("N%03X%03X", n.Y, n.X)
				graph.AddWeightedEdge(node, neighborNode, (*g)[n.Y][n.X])
			}
		}
	}

	start := "N000000"
	end := fmt.Sprintf("N%03X%03X", len(*g)-1, len((*g)[len(*g)-1])-1)
	distances := graph.Dijkstra(start)
	fmt.Println(start, "->", end, ":", distances[end])
}

func part2(g *grid) {
	newGrid := extendGrid(g)
	part1(newGrid)
}

func extendGrid(g *grid) *grid {
	N := 5
	newGrid := make([][]int, N*(len(*g)))
	for i := 0; i < len(newGrid); i++ {
		newGrid[i] = make([]int, N*len((*g)[0]))
	}

	for i := 0; i < len(*g); i++ {
		for j := 0; j < len((*g)[i]); j++ {
			newGrid[i][j] = (*g)[i][j]
		}
	}

	X := len(*g)

	// fill columns
	for repeat := 1; repeat < N; repeat++ {
		for i := 0; i < X; i++ { // X rows down
			for j := 0; j < len((*g)[0]); j++ {
				u := newGrid[i][(repeat-1)*X+j]
				v := u + 1
				if v > 9 {
					v = 1
				}
				newGrid[i][repeat*X+j] = v
			}
		}
	}

	// fill rows
	for repeat := 1; repeat < N; repeat++ {
		for i := 0; i < X; i++ { // X rows
			for j := 0; j < len(newGrid[0]); j++ {
				u := newGrid[(repeat-1)*X+i][j]
				v := u + 1
				if v > 9 {
					v = 1
				}
				newGrid[repeat*X+i][j] = v
			}
		}
	}

	return (*grid)(unsafe.Pointer(&newGrid))
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
	//part1((*grid)(unsafe.Pointer(&g1)))
	part2((*grid)(unsafe.Pointer(&g2)))
}
