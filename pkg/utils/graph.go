package utils

import "math"

type Graph struct {
	Edges    map[string]map[string]int
	Directed bool
}

func New() *Graph {
	return &Graph{}
}

func (g *Graph) AddVertex(v string) {
	if g.Edges == nil {
		g.Edges = make(map[string]map[string]int)
	}

	if _, ok := g.Edges[v]; !ok {
		g.Edges[v] = make(map[string]int)
	}
}

func (g *Graph) VertexCount() int {
	return len(g.Edges)
}

func (g *Graph) AddEdge(from, to string) {
	g.AddWeightedEdge(from, to, 0)
}

func (g *Graph) AddWeightedEdge(from, to string, weight int) {
	g.AddVertex(from)
	g.AddVertex(to)

	g.Edges[from][to] = weight
	if !g.Directed {
		g.Edges[to][from] = weight
	}
}

func OkToVisitOnceOnly(v string, visited map[string]int) bool {
	if _, ok := visited[v]; ok && visited[v] > 0 {
		return false
	}
	return true
}

// https://www.codingame.com/playgrounds/1608/shortest-paths-with-dijkstras-algorithm/ending
func (g *Graph) Dijkstra(start string) map[string]int {
	dist := map[string]int{}
	visited := map[string]bool{}
	for k, _ := range g.Edges {
		dist[k] = math.MaxInt
		visited[k] = false
	}

	dist[start] = 0

	mindistance := func(d map[string]int, visited map[string]bool) string {
		minDistance := math.MaxInt
		minV := ""
		for v, distance := range d {
			if distance < minDistance && !visited[v] {
				minV = v
				minDistance = distance
			}
		}
		return minV
	}

	C := start
	for C != "" {
		for N, distance := range g.Edges[C] {
			x := dist[C] + distance
			if x < dist[N] {
				dist[N] = x
			}
		}
		visited[C] = true
		C = mindistance(dist, visited)
	}

	return dist
}

func (g *Graph) DFS(start, end string, okToVisit func(string, map[string]int) bool, visited map[string]int, currentPath []string, paths *[][]string) {
	if !okToVisit(start, visited) {
		return
	}

	visited[start]++
	currentPath = append(currentPath, start)

	// we reached the end
	if start == end {
		*paths = append(*paths, currentPath)
		visited[start]--
		currentPath = currentPath[:len(currentPath)-1]
		return
	}

	for next := range g.Edges[start] {
		g.DFS(next, end, okToVisit, visited, currentPath, paths)
	}

	currentPath = currentPath[:len(currentPath)-1]
	visited[start]--
}

func (g *Graph) AllPaths(start, end string, okToVisit func(string, map[string]int) bool) [][]string {
	paths := [][]string{}
	g.DFS(start, end, okToVisit, map[string]int{}, []string{}, &paths)
	return paths
}
