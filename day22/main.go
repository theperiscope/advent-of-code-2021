package main

import (
	"AOC/pkg/utils"
	"fmt"
	"os"
)

// https://github.com/akosgarai/coldet/blob/master/coldet.go
// https://developer.mozilla.org/en-US/docs/Games/Techniques/3D_collision_detection
// Axis aligned bounding box
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
	op   string
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
		//fmt.Printf("%s x=%d..%d,y=%d..%d,z=%d..%d\n", op, x1, x2, y1, y2, z1, z2)
		c := cuboid{op, AABB{origin: point3D{x1, y1, z1}, width: x2 - x1 + 1, height: y2 - y1 + 1, length: z2 - z1 + 1}}
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
	c1 := c.origin
	c2 := point3D{c1.X + c.width, c1.Y + c.height, c1.Z + c.length}
	o1 := other.origin
	o2 := point3D{o1.X + other.width, o1.Y + other.height, o1.Z + other.length}

	intXFrom := utils.MaxInt(c1.X, o1.X)
	intXTo := utils.MinInt(c2.X, o2.X)
	intYFrom := utils.MaxInt(c1.Y, o1.Y)
	intYTo := utils.MinInt(c2.Y, o2.Y)
	intZFrom := utils.MaxInt(c1.Z, o1.Z)
	intZTo := utils.MinInt(c2.Z, o2.Z)

	if intXFrom > intXTo || intYFrom > intYTo || intZFrom > intZTo {
		return nil
	}

	return &AABB{origin: point3D{X: intXFrom, Y: intYFrom, Z: intZFrom}, width: intXTo - intXFrom, height: intYTo - intYFrom, length: intZTo - intZFrom}
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

	newCuboids := resolveIntersections(initializationCuboids)
	fmt.Println("Cubes Lit:", countLit(newCuboids))
}

func part2(cuboids []cuboid) {
	newCuboids := resolveIntersections(cuboids)
	fmt.Println("Cubes Lit:", countLit(newCuboids))
}

func countLit(cuboids []cuboid) (cnt int) {
	for _, cuboid := range cuboids {
		if cuboid.op == "on" {
			cnt += cuboid.aabb.volume()
		} else {
			cnt -= cuboid.aabb.volume()
		}
	}
	return
}

func resolveIntersections(cuboids []cuboid) (result []cuboid) {
	result = []cuboid{{aabb: cuboids[0].aabb, op: "on"}} // the first cuboid is lit in all datasets
	for i := 1; i < len(cuboids); i++ {
		for _, r := range result {
			// intersection of 2 cuboids is always 1 cuboid
			overlapAABB := r.aabb.overlap(cuboids[i].aabb)
			if overlapAABB != nil {
				var newOp string
				if r.op == "on" {
					newOp = "off"
				} else {
					newOp = "on"
				}
				overlapCuboid := cuboid{aabb: *overlapAABB, op: newOp}
				result = append(result, overlapCuboid)
			}
		}
		if cuboids[i].op == "on" {
			result = append(result, cuboid{aabb: cuboids[i].aabb, op: "on"})
		}
	}
	return
}
