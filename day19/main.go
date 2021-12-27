package main

import (
	"AOC/pkg/utils"
	"fmt"
	"os"
	"strings"
)

type point3D struct {
	X, Y, Z int
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

func calculateBestOverlap(s0, s1 []point3D) (point3D, int, matrix3D) {
	// computer matrixes for the 24 unique orientations, per https://stackoverflow.com/questions/16452383/how-to-get-all-24-rotations-of-a-3-dimensional-array
	orientations := []string{"X", "Y", "Z", "XX", "XY", "XZ", "YX", "YY", "ZY", "ZZ", "XXX", "XXY", "XXZ", "XYX", "XYY", "XZZ", "YXX", "YYY", "ZZZ", "XXXY", "XXYX", "XYXX", "XYYY"}
	orientationMatrices := []matrix3D{{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}} // start with ident
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
		orientationMatrices = append(orientationMatrices, m)
	}

	optimumTransform := matrix3D{}
	optimumOffset := point3D{}
	optimumOverlaps := 0
	for _, m := range orientationMatrices {
		transformedS1 := []point3D{}
		for _, x := range s1 {
			transformedS1 = append(transformedS1, x.Multiply(m))
		}

		for i, u := range s0 {
			for j, v := range transformedS1 {
				if j > i {
					continue
				}

				offset := v.Subtract(u)
				offsetTransformedS1 := []point3D{}
				for _, x := range transformedS1 {
					offsetTransformedS1 = append(offsetTransformedS1, x.Subtract(offset))
				}

				overlaps := 0
				for _, x := range s0 {
					for _, y := range offsetTransformedS1 {
						if x == y {
							overlaps++
						}
					}
				}

				if overlaps > optimumOverlaps {
					optimumOverlaps = overlaps
					optimumOffset = offset
					optimumTransform = m
					if overlaps >= 12 {
						return optimumOffset, optimumOverlaps, optimumTransform
					}
				}
			}
		}
	}
	return optimumOffset, optimumOverlaps, optimumTransform
}

func uniquePoints(points []point3D) []point3D {
	seen := make(map[point3D]bool)
	newPoints := []point3D{}
	for _, p := range points {
		if _, value := seen[p]; !value {
			seen[p] = true
			newPoints = append(newPoints, p)
		}
	}
	return newPoints
}

func part1(lines []string) (scannerLocations []point3D, alignedPointsCount int) {
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

	alignedPoints := scanners[0]
	remainingScanners := scanners[1:]
	scannerLocations = []point3D{{0, 0, 0}}
	fmt.Printf("Processed %d/%d, %d aligned points.\n", len(scannerLocations), len(scanners), len(alignedPoints))

	for len(remainingScanners) > 0 {
		for i, current := range remainingScanners {
			scannerLocation, overlap, transform := calculateBestOverlap(alignedPoints, current)
			if overlap < 12 {
				continue
			}
			realignedCurrentPoints := []point3D{}
			for _, x := range current {
				realignedCurrentPoints = append(realignedCurrentPoints, x.Multiply(transform).Subtract(scannerLocation))
			}
			alignedPoints = uniquePoints(append(alignedPoints, realignedCurrentPoints...))
			remainingScanners = append(remainingScanners[:i], remainingScanners[i+1:]...) // cut out current
			fmt.Printf("Processed %d/%d, %d aligned points.\n", len(scannerLocations), len(scanners), len(alignedPoints))
			scannerLocations = append(scannerLocations, scannerLocation)
			break
		}
	}

	return scannerLocations, len(alignedPoints)
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
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Printf("Usage: %s <inputfile>\n", utils.GetProgramName())
		return
	}

	lines, _ := utils.ReadInput(argsWithoutProg[0])

	scannerLocations, alignedPointsCount := part1(lines)
	fmt.Println("Part 1 Answer:", alignedPointsCount)

	maxDistance := part2(scannerLocations)
	fmt.Println("Part 2 Answer:", maxDistance)
}
