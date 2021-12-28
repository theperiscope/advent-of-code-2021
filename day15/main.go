package main

import (
	"AOC/pkg/utils"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
	"unsafe"
)

type grid [][]int

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

type position struct{ x, y int }
type path struct {
	location position
	cost     int
}

func (g *grid) distances(x, y int) (m []path) {
	m = []path{}
	for _, n := range [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
		if x+n[1] < 0 || x+n[1] > len((*g)[0])-1 || y+n[0] < 0 || y+n[0] > len((*g))-1 { // valid range check
			continue
		}
		m = append(m, path{location: position{x: x + n[1], y: y + n[0]}, cost: (*g)[y+n[0]][x+n[1]]})
	}
	return m
}

func (g *grid) solve(start position, end position) int {
	paths := pathPriorityQueue{}
	bestCost := math.MaxInt
	visited := map[position]bool{}
	heap.Init(&paths)
	var p *path = &path{location: start, cost: 0}
	for p != nil {
		nextMoves := g.distances(p.location.x, p.location.y)
		for _, move := range nextMoves {
			next := path{move.location, p.cost + move.cost}
			if move.location == end {
				if next.cost < bestCost {
					bestCost = next.cost
				}
			} else {
				heap.Push(&paths, &next)
			}
		}
		visited[p.location] = true
		p = nil
		for len(paths) > 0 {
			p = heap.Pop(&paths).(*path)
			if visited[p.location] { // skip already visited locations
				p = nil
				continue
			}
			break
		}
		if p == nil {
			break
		}
	}
	return bestCost
}

func part1(g *grid) {
	startTime := time.Now()
	bestCost := g.solve(position{0, 0}, position{x: len((*g)[0]) - 1, y: len(*g) - 1})
	duration := time.Since(startTime)
	fmt.Println("Part 1 Answer:", bestCost, "in", duration)
}

func part2(g *grid) {
	newGrid := extendGrid(g)
	startTime := time.Now()
	bestCost := newGrid.solve(position{0, 0}, position{x: len((*newGrid)[0]) - 1, y: len(*newGrid) - 1})
	duration := time.Since(startTime)
	fmt.Println("Part 2 Answer:", bestCost, "in", duration)
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
	part1((*grid)(unsafe.Pointer(&g1)))
	part2((*grid)(unsafe.Pointer(&g2)))
}

// A pathPriorityQueue implements heap.Interface and holds Items.
type pathPriorityQueue []*path

func (pq pathPriorityQueue) Len() int { return len(pq) }

func (pq pathPriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq pathPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pathPriorityQueue) Push(x interface{}) {
	item := x.(*path)
	*pq = append(*pq, item)
}

func (pq *pathPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
