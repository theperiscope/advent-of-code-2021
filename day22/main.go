package main

import (
	"AOC/pkg/utils"
	"fmt"
	"os"
)

// https://github.com/akosgarai/coldet/blob/master/coldet.go
// https://developer.mozilla.org/en-US/docs/Games/Techniques/3D_collision_detection
// AABB represents axis-aligned bounding box
type AABB struct {
	origin point3D
	width  int // X axis
	height int // Y axis
	length int // Z axis
}

type point3D struct {
	X, Y, Z int
}

type cuboid struct {
	lit  bool
	aabb AABB
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Printf("Usage: %s <inputfile>\n", utils.GetProgramName())
		return
	}

	lines, _ := utils.ReadInput(argsWithoutProg[0])
	cuboids := []cuboid{}
	for _, line := range lines {
		var op string
		var x1, x2, y1, y2, z1, z2 int
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &op, &x1, &x2, &y1, &y2, &z1, &z2)
		c := cuboid{op == "on", AABB{origin: point3D{x1, y1, z1}, width: x2 - x1 + 1, height: y2 - y1 + 1, length: z2 - z1 + 1}}
		cuboids = append(cuboids, c)
	}

	part1(cuboids)
	part2(cuboids)
}

func (c AABB) isFullyContainedIn(other AABB) bool {
	return c.origin.X >= other.origin.X && c.origin.X+c.width <= other.width &&
		c.origin.Y >= other.origin.Y && c.origin.Y+c.height <= other.height &&
		c.origin.Z >= other.origin.Z && c.origin.Z+c.length <= other.length
}

func (c AABB) overlap(other AABB) *AABB {
	cMin, cMax := c.origin, point3D{c.origin.X + c.width, c.origin.Y + c.height, c.origin.Z + c.length}
	oMin, oMax := other.origin, point3D{other.origin.X + other.width, other.origin.Y + other.height, other.origin.Z + other.length}

	overlapXFrom, overlapXTo := utils.MaxInt(cMin.X, oMin.X), utils.MinInt(cMax.X, oMax.X)
	overlapYFrom, overlapYTo := utils.MaxInt(cMin.Y, oMin.Y), utils.MinInt(cMax.Y, oMax.Y)
	overlapZFrom, overlapZTo := utils.MaxInt(cMin.Z, oMin.Z), utils.MinInt(cMax.Z, oMax.Z)

	if overlapXFrom > overlapXTo || overlapYFrom > overlapYTo || overlapZFrom > overlapZTo { // invalid overlap bounds
		return nil
	}
	return &AABB{origin: point3D{X: overlapXFrom, Y: overlapYFrom, Z: overlapZFrom}, width: overlapXTo - overlapXFrom, height: overlapYTo - overlapYFrom, length: overlapZTo - overlapZFrom}
}

func (c AABB) volume() int {
	return c.width * c.height * c.length
}

func part1(cuboids []cuboid) {
	initializationCuboids := []cuboid{}
	for _, c := range cuboids {
		if c.aabb.isFullyContainedIn(AABB{origin: point3D{-50, -50, -50}, width: 100, height: 100, length: 100}) {
			initializationCuboids = append(initializationCuboids, c)
		}
	}
	fmt.Println("Part 1 Answer:", countLit(initializationCuboids))
}

func part2(cuboids []cuboid) {
	fmt.Println("Part 2 Answer:", countLit(cuboids))
}

func countLit(cuboids []cuboid) (count int) {
	resolveIntersections := func(cuboids []cuboid) (resolved []cuboid) {
		resolved = []cuboid{{aabb: cuboids[0].aabb, lit: true}} // the first cuboid is lit in all datasets
		for i := 1; i < len(cuboids); i++ {
			current, newOverlaps := cuboids[i], []cuboid{}
			for _, r := range resolved {
				overlapAABB := r.aabb.overlap(current.aabb) // overlap of 2 AABB cuboids is always 1 AABB cuboid
				if overlapAABB != nil {
					newOverlaps = append(newOverlaps, cuboid{aabb: *overlapAABB, lit: !r.lit})
				}
			}
			if len(newOverlaps) > 0 { // add the overlaps
				resolved = append(resolved, newOverlaps...)
			}
			if current.lit { // then, add the current cuboid if lit
				resolved = append(resolved, current)
			}
		}
		return
	}

	newCuboids := resolveIntersections(cuboids)
	for _, cuboid := range newCuboids {
		if cuboid.lit {
			count += cuboid.aabb.volume()
		} else {
			count -= cuboid.aabb.volume()
		}
	}
	return
}
