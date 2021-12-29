package main

import (
	"AOC/pkg/utils"
	"fmt"
	"log"
	"strings"
)

const EMPTY string = "."

type grid [][]string
type path struct {
	grid *grid
	cost int
}

type position struct{ x, y int }
type move struct {
	location position
	dist     int
}

func parseInput(lines []string) *grid {
	g := grid{}
	for _, line := range lines {
		g = append(g, strings.Split(line, ""))
	}
	return &g
}

func moveCost(s string) int {
	return map[string]int{"A": 1, "B": 10, "C": 100, "D": 1000}[s]
}

func targetX(s string) int {
	return map[string]int{"A": 3, "B": 5, "C": 7, "D": 9}[s]
}

func (g *grid) uniqueStateKey() string {
	ss := ""
	for y := 1; y < len(*g)-1; y++ {
		ss += strings.Replace(strings.Join((*g)[y], ""), "#", "", -1)
	}
	return ss
}

func (g *grid) clone() *grid {
	gg := grid{}
	for y := 0; y < len(*g); y++ {
		gg = append(gg, []string{})
		gg[y] = make([]string, len((*g)[0]))
		for x := 0; x < len((*g)[0]); x++ {
			gg[y][x] = (*g)[y][x]
		}
	}
	return &gg
}

func (g *grid) distanceMap(x, y int) []move {
	var possibleMoves func(g *grid, x, y, dist int) []move
	possibleMoves = func(g *grid, x, y, dist int) []move {
		m := []move{}
		if (*g)[y][x] == EMPTY {
			(*g)[y][x] = fmt.Sprintf("%d", dist)
			m = append(m, move{location: position{x, y}, dist: dist})
		}
		for _, n := range [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
			if (*g)[y+n[0]][x+n[1]] == EMPTY {
				m = append(m, possibleMoves(g, x+n[1], y+n[0], dist+1)...)
			}
		}
		return m
	}
	return possibleMoves(g.clone(), x, y, 0)
}

func (g *grid) isAmphipodsOwnHouse(destX, destY, startX, startY int) bool {
	return (destY > 1) && destX == targetX((*g)[startY][startX])
}

func (g *grid) isAmphipodsHouseClean(startX, startY int) bool {
	rows, targetX := len(*g)-1, targetX((*g)[startY][startX])
	for y := 2; y < rows-1; y++ {
		if !strings.Contains(EMPTY+(*g)[startY][startX], (*g)[y][targetX]) {
			return false
		}
	}
	return true
}

func (g *grid) nextMoves(startX, startY, rows, cols int) []move {
	distancesTo, moves := g.clone().distanceMap(startX, startY), []move{}
	for _, entry := range distancesTo { // check distance map entries against our rules
		destX, destY := entry.location.x, entry.location.y
		if (startY == 1 && destY == 1) || (startY != 1 && destY != 1) || (startY != 1 && (destY == 1 && (destX == 3 || destX == 5 || destX == 7 || destX == 9))) {
			continue
		}
		if startY == 1 && (!g.isAmphipodsOwnHouse(destX, destY, startX, startY) || !g.isAmphipodsHouseClean(startX, startY)) {
			continue
		}
		if startY == 1 && g.isAmphipodsOwnHouse(destX, destY, startX, startY) && g.isAmphipodsHouseClean(startX, startY) {
			if !strings.Contains("#"+(*g)[startY][startX], (*g)[destY+1][destX]) { // we can only move to deepest open spot in house
				continue
			}
		}
		moves = append(moves, entry)
	}
	return moves
}

func (p *path) findAmphipods(rows, cols int) []position {
	amphipods := []position{}
	for y := 1; y < rows-1; y++ {
		for x := 1; x < cols-1; x++ {
			if !strings.Contains("ABCD", (*p.grid)[y][x]) {
				continue
			}

			ch, isRoomSet := (*p.grid)[y][x], false
			if x == targetX(ch) {
				isRoomSet = true
				for j := y + 1; j < rows-1; j++ {
					if (*p.grid)[j][x] != ch {
						isRoomSet = false
						break
					}
				}
				if isRoomSet {
					continue
				}
			}

			amphipods = append(amphipods, position{x, y})
		}
	}
	return amphipods
}

func solve(start, end []string) []path {
	startGrid, endGrid := parseInput(start), parseInput(end)
	rows, cols := len(*startGrid), len((*startGrid)[0])
	paths, completePaths := []*path{{startGrid, 0}}, []path{}
	bestStates := map[string]int{}

	for len(paths) > 0 {
		p := paths[len(paths)-1]
		paths = paths[0 : len(paths)-1]
		amphipods := p.findAmphipods(rows, cols)
		for _, amphipod := range amphipods {
			x, y := amphipod.x, amphipod.y
			ch, nextMoves := (*p.grid)[y][x], p.grid.nextMoves(x, y, rows, cols)
			for _, move := range nextMoves {
				next := path{}
				next.grid = p.grid.clone()
				(*next.grid)[y][x] = EMPTY
				(*next.grid)[move.location.y][move.location.x] = ch
				k := next.grid.uniqueStateKey()
				next.cost = p.cost + moveCost(ch)*move.dist
				if _, ok := bestStates[k]; !ok || next.cost < bestStates[k] {
					bestStates[k] = next.cost
					if k == endGrid.uniqueStateKey() {
						completePaths = append(completePaths, next)
					} else {
						paths = append(paths, &next)
					}
				}
			}
		}
	}
	return completePaths
}

func part1() {
	lines, err := utils.ReadInput("input-part1.txt")
	if err != nil {
		log.Fatal(err)
	}
	start := lines[0:5]
	end := lines[6:11]

	finalPaths := solve(start, end)
	fmt.Printf("Part 1 Answer: %d\n", finalPaths[len(finalPaths)-1].cost)
}

func part2() {
	lines, err := utils.ReadInput("input-part2.txt")
	if err != nil {
		log.Fatal(err)
	}
	start := lines[0:7]
	end := lines[8:15]

	finalPaths := solve(start, end)
	fmt.Printf("Part 2 Answer: %d\n", finalPaths[len(finalPaths)-1].cost)
}

func main() {
	part1()
	part2()
}
