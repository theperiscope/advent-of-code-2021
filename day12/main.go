package main

import (
	"AOC/pkg/utils"
	graph "AOC/pkg/utils"
	"fmt"
	"os"
	"strings"
)

func part1(lines []string) {
	// build graph
	g := graph.New()
	for _, line := range lines {
		edge := strings.Split(line, "-")
		g.AddEdge(edge[0], edge[1])
	}

	okToVisitUnvisitedAndLargeCaves := func(u string, visited map[string]int) bool {
		if _, ok := visited[u]; ok {
			if visited[u] > 0 && utils.IsLower(u) { // small cave
				return false
			}
		}
		return true
	}

	allPaths := g.AllPaths("start", "end", okToVisitUnvisitedAndLargeCaves)
	fmt.Println("Part 1 Answer:", len(allPaths))
}

func part2(lines []string) {
	// build graph
	g := graph.New()
	for _, line := range lines {
		edge := strings.Split(line, "-")
		g.AddEdge(edge[0], edge[1])
	}

	// identify all small caves
	smallCaves := []string{}
	for u, _ := range g.Edges {
		if utils.IsLower(u) && u != "start" && u != "end" {
			smallCaves = append(smallCaves, u)
		}
	}

	okToVisitUnvisitedAndLargeCavesAndOneSmallCaveTwice := func(u string, visited map[string]int) bool {
		if _, ok := visited[u]; ok {
			if u == "start" { // start is an exception
				return false
			}
			if visited[u] > 0 && utils.IsLower(u) { // small cave

				canVisitTwice := true
				for _, cave := range smallCaves {
					_, ok := visited[cave]
					canVisitTwice = canVisitTwice && (!ok || visited[cave] <= 1)
					if !canVisitTwice {
						return false
					}
				}
			}
		}
		return true
	}

	allPaths := g.AllPaths("start", "end", okToVisitUnvisitedAndLargeCavesAndOneSmallCaveTwice)
	fmt.Println("Part 2 Answer:", len(allPaths))
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Printf("Usage: %s <inputfile>\n", utils.GetProgramName())
		return
	}

	lines, _ := utils.ReadInput(argsWithoutProg[0])
	part1(lines)
	part2(lines)
}
