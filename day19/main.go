package main

import (
	"AOC/pkg/utils"
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

type point3D struct {
	X, Y, Z int
}

func (p point3D) Add(other point3D) point3D {
	return point3D{p.X + other.X, p.Y + other.Y, p.Z + other.Z}
}
func (p point3D) Subtract(other point3D) point3D {
	return point3D{p.X - other.X, p.Y - other.Y, p.Z - other.Z}
}
func (p point3D) ManhattanDistance(other point3D) int {
	return utils.AbsInt(p.X-other.X) + utils.AbsInt(p.Y-other.Y) + utils.AbsInt(p.Z-other.Z)
}
func (p point3D) Multiply(matrix matrix3D) point3D {
	return point3D{
		p.X*matrix[0][0] + p.Y*matrix[0][1] + p.Z*matrix[0][2],
		p.X*matrix[1][0] + p.Y*matrix[1][1] + p.Z*matrix[1][2],
		p.X*matrix[2][0] + p.Y*matrix[2][1] + p.Z*matrix[2][2],
	}
}
func (a point3D) EucledanDistance(b point3D) int {
	x := math.Pow(float64(a.X)-float64(b.X), 2)
	y := math.Pow(float64(a.Y)-float64(b.Y), 2)
	z := math.Pow(float64(a.Z)-float64(b.Z), 2)
	return int(x + y + z)
}

type matrix3D [3][3]int

func (m matrix3D) Multiply(other matrix3D) matrix3D {
	result := matrix3D{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			result[i][j] = m[i][0]*other[0][j] + m[i][1]*other[1][j] + m[i][2]*other[2][j]
		}
	}
	return result
}

func part1(lines []string) (alignedPoints map[point3D]bool, scannerLocations []point3D) {
	var scanners [][]point3D
	i := -1
	for _, line := range lines {
		if strings.HasPrefix(line, "---") {
			i++
			scanners = append(scanners, []point3D{})
			continue
		}
		if line == "" {
			continue
		}
		var x, y, z int
		_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			panic(err)
		}
		scanners[i] = append(scanners[i], point3D{x, y, z})
	}

	orientations := []string{"X", "Y", "Z", "XX", "XY", "XZ", "YX", "YY", "ZY", "ZZ", "XXX", "XXY", "XXZ", "XYX", "XYY", "XZZ", "YXX", "YYY", "ZZZ", "XXXY", "XXYX", "XYXX", "XYYY"}
	orientationMatrices := map[string]matrix3D{}
	for _, orientation := range orientations {
		m := matrix3D{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
		for _, c := range orientation {
			if c == 'X' {
				m = m.Multiply(matrix3D{{0, -1, 0}, {1, 0, 0}, {0, 0, 1}})
			} else if c == 'Y' {
				m = m.Multiply(matrix3D{{0, 0, -1}, {0, 1, 0}, {1, 0, 0}})
			} else if c == 'Z' {
				m = m.Multiply(matrix3D{{1, 0, 0}, {0, 0, -1}, {0, 1, 0}})
			}
		}
		orientationMatrices[orientation] = m
	}
	orientationMatrices["IDENT"] = matrix3D{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

	alignedPoints = map[point3D]bool{}
	for _, p := range scanners[0] {
		alignedPoints[p] = true
	}
	remainingScanners := scanners[1:]
	scannerLocations = []point3D{{0, 0, 0}}
	fmt.Printf("Processed %d/%d, %d aligned points.\r", len(scannerLocations), len(scannerLocations)+len(remainingScanners), len(alignedPoints))

	for len(remainingScanners) > 0 {
		for i := 0; i < len(remainingScanners); i++ {
			current := remainingScanners[i]
			bestMatches, bestTransform, scannerLocation := alignScanners(alignedPoints, current, orientationMatrices)

			if bestMatches < 12 {
				continue
			}

			scannerLocations = append(scannerLocations, scannerLocation)
			for _, v := range current {
				p := v.Multiply(bestTransform).Add(scannerLocation)
				alignedPoints[p] = true
			}
			remainingScanners, i = append(remainingScanners[:i], remainingScanners[i+1:]...), 0 // cut out current and start over
			fmt.Printf("Processed %d/%d, %d aligned points.\r", len(scannerLocations), len(scanners), len(alignedPoints))
		}
	}
	fmt.Println()
	return
}

func part2(scannerLocations []point3D) (maxDistance int) {
	maxDistance = 0
	for i, a := range scannerLocations {
		for j, b := range scannerLocations {
			if j >= i {
				break
			}
			d := a.ManhattanDistance(b)
			if d > maxDistance {
				maxDistance = d
			}
		}
	}

	return maxDistance
}

func main() {
	utils.AssertArgs()
	lines := utils.AssertInput()

	alignedPoints, scannerLocations := part1(lines)
	fmt.Println("Part 1 Answer:", len(alignedPoints))

	maxDistance := part2(scannerLocations)
	fmt.Println("Part 2 Answer:", maxDistance)
}

func alignScanners(alignedPointsMap map[point3D]bool, currentScanner []point3D, orientationMatrices map[string]matrix3D) (int, matrix3D, point3D) {
	alignedPoints := []point3D{}
	for p, _ := range alignedPointsMap {
		alignedPoints = append(alignedPoints, p)
	}
	m0, m1 := distanceMatrix(alignedPoints), distanceMatrix(currentScanner)
	sharedPoints := findSharedPoints(m0, m1)
	if len(sharedPoints) < 12 {
		return 0, matrix3D{}, point3D{}
	}
	// check which orientation matrix will match with currently aligned points
	// idea for future improvements is to evaluate use of SVD or Umeyama's algorithm instead http://nghiaho.com/?page_id=671
	for _, m := range orientationMatrices {
		offset := point3D{}
		for k, v := range sharedPoints {
			s0p, s1p := alignedPoints[k], currentScanner[v].Multiply(m)
			offset = s0p.Subtract(s1p) // first offset should apply for all points if it is the correct scanner location
			break
		}
		n := 0
		for k, v := range sharedPoints {
			s0p, s1p := alignedPoints[k], currentScanner[v].Multiply(m).Add(offset)
			if s1p == s0p { // we expect here all to match if we have the correct scanner location
				n++
			} else {
				break // it is an incorrect orientation matrix
			}
		}
		if n >= 12 {
			return n, m, offset
		}
	}
	return 0, matrix3D{}, point3D{}
}

func distanceMatrix(points []point3D) [][]int {
	result := make([][]int, len(points))

	for i := 0; i < len(points); i++ {
		result[i] = make([]int, len(points))
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			result[i][j] = points[i].EucledanDistance(points[j])
			result[j][i] = result[i][j]
		}
	}
	return result
}

type workItem struct {
	rowIndex1, rowIndex2       int
	distanceMap1, distanceMap2 [][]int
}

// findSharedPoints returns map of points from distanceMap1 which are the equivalent in distanceMap2
func findSharedPoints(distanceMap1, distanceMap2 [][]int) map[int]int {
	var doWork = func(rowIndex1, rowIndex2 int, dmap1, dmap2 [][]int) map[int]int {
		row1, row2 := dmap1[rowIndex1], dmap2[rowIndex2]
		potentialMatches := inAAndB(row1, row2)
		if len(potentialMatches) >= 12 {
			m := map[int]int{}
			for col1 := 0; col1 < len(dmap1); col1++ {
				for col2 := 0; col2 < len(dmap2); col2++ {
					if dmap1[rowIndex1][col1] == dmap2[rowIndex2][col2] {
						m[rowIndex1], m[col1] = rowIndex2, col2
					}
				}
			}
			return m
		}
		return map[int]int{}
	}

	parallelization := 16
	var wg sync.WaitGroup
	wg.Add(parallelization)
	c := make(chan workItem)
	result := map[int]int{}
	for i := 0; i < parallelization; i++ {
		go func(c chan workItem) {
			for {
				v, more := <-c
				if !more {
					wg.Done()
					return
				}

				vv := doWork(v.rowIndex1, v.rowIndex2, v.distanceMap1, v.distanceMap2)
				if len(vv) >= 12 {
					result = vv
				}
			}
		}(c)
	}
	// found during testing that breaking up processing into 3 separate loops makes
	// performance better than going in-sequence one-by-one
	for x1 := 0; x1 < len(distanceMap1); x1 += 3 { // %3 == 0
		for x2 := 0; x2 < len(distanceMap2); x2 += 3 {
			c <- workItem{x1, x2, distanceMap1, distanceMap2}
			if len(result) > 0 { // if we have result, no need to queue up anything else
				break
			}
		}
	}
	if len(result) == 0 { // still nothing found, next batch
		for x1 := 1; x1 < len(distanceMap1); x1 += 3 { // %3 == 1
			for x2 := 1; x2 < len(distanceMap2); x2 += 3 {
				c <- workItem{x1, x2, distanceMap1, distanceMap2}
				if len(result) > 0 { // if we have result, no need to queue up anything else
					break
				}
			}
		}
	}
	if len(result) == 0 { // still nothing found, final batch
		for x1 := 2; x1 < len(distanceMap1); x1 += 3 { // %3 == 2
			for x2 := 2; x2 < len(distanceMap2); x2 += 3 {
				c <- workItem{x1, x2, distanceMap1, distanceMap2}
				if len(result) > 0 { // if we have result, no need to queue up anything else
					break
				}
			}
		}
	}
	close(c)
	wg.Wait()
	return result
}

// inAAndB interestion implementation using sorting
func inAAndB(aa, bb []int) (result []int) {
	if len(aa) == 0 || len(bb) == 0 {
		return []int{}
	}
	a, b := make([]int, len(aa)), make([]int, len(bb))
	copy(a, aa)
	copy(b, bb)
	sort.Ints(a)
	sort.Ints(b)
	for i, j := 0, 0; i < len(a) && j < len(b); {
		x, y := a[i], b[j]
		if x == y {
			if len(result) == 0 || x > result[len(result)-1] {
				result = append(result, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return result
}
