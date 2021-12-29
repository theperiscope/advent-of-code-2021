package main

import (
	"AOC/pkg/utils"
	"fmt"
	"math"
)

type point struct {
	x int
	y int
}

func (p point) String() string {
	return fmt.Sprintf("(X=%d, Y=%d)", p.x, p.y)
}

type bounds struct {
	topLeft     point
	bottomRight point
}

func (b bounds) Contains(p point) bool {
	return p.x >= b.topLeft.x && p.x <= b.bottomRight.x && p.y <= b.topLeft.y && p.y >= b.bottomRight.y
}

func part1(start point, target bounds) {
	// > Due to drag, the probe's x velocity changes by 1 toward the value 0; that is, it decreases by 1 if it is greater than 0, increases by 1
	// > if it is less than 0, or does not change if it is already 0.
	//
	// * for probe launcher X-position to be in target area, we need to move from start X to somewhere between TopLeft.X and BottomRight.X
	// * whether positive or negative for selected initial speed Δx  we'll move Δx + Δx-1 + Δx-2 + .. + 1  => SUM(1..Δx) which is Δx*(Δx-1)*2
	minX, maxX := float64(utils.AbsInt(target.topLeft.x)), float64(utils.AbsInt(target.bottomRight.x))
	// n*n + n - 2*sum(1..n) = 0
	a, b, cMinX, cMaxX := float64(1), float64(1), float64(-minX*2), float64(-maxX*2)
	// we only need the positive solution to calculate optimum range for X
	startX := int(math.Floor((-b + math.Sqrt(math.Pow(b, 2)-(4*a*cMinX))) / (2 * a)))
	endX := int(math.Ceil((-b + math.Sqrt(math.Pow(b, 2)-(4*a*cMaxX))) / (2 * a)))

	// for optimal Y, we pick -BottomRight.Y - 1
	// for range startX..endX, most "y" is gained if we picked opposite of bottom-y minus one.
	// essetially, y increases as X approaches target area and Δx approaches 0. once there, Δx is 0 and Y
	// starts decreasing until we hit the target area. If we picked just BottomRight.Y we will overshoot target area
	ΔyOptimum := -target.bottomRight.y - 1

	ΔxOptimum, maxY := math.MinInt, math.MinInt
	for Δx := startX; Δx <= endX; Δx++ {
		maxYCurrent := shoot(target, start, Δx, ΔyOptimum)
		if maxYCurrent > maxY {
			maxY = maxYCurrent
			ΔxOptimum = Δx
		}
	}
	fmt.Printf("Part 1 Answer: MaxY=%d for Δx=%d, Δy=%d\n", maxY, ΔxOptimum, ΔyOptimum)
}

func part2(start point, target bounds) {
	// brute-force it, perhaps better numbers can be picked
	n, ΔxMin, ΔxMax, ΔyMin, ΔyMax := 0, 0, utils.AbsInt(target.bottomRight.x), -utils.AbsInt(target.bottomRight.y), utils.AbsInt(target.bottomRight.y)
	for Δx := ΔxMin; Δx <= ΔxMax; Δx++ {
		for Δy := ΔyMin; Δy <= ΔyMax; Δy++ {
			maxY := shoot(target, start, Δx, Δy)
			if maxY != math.MinInt {
				n++
			}
		}
	}
	fmt.Printf("Part 2 Answer: %d\n", n)
}

func shoot(target bounds, p point, Δx, Δy int) int {
	next, maxY := p, math.MinInt
	for {
		if next.y > maxY {
			maxY = next.y
		}
		if target.Contains(next) {
			return maxY
		}
		next = point{x: next.x + Δx, y: next.y + Δy}
		if (next.x < target.topLeft.x || next.x > target.bottomRight.x) && Δx == 0 {
			break // X is before/after target with dx = 0, STOP
		}
		if next.y < target.bottomRight.y {
			break // Y will only keep decreasing, STOP
		}
		if Δx > 0 {
			Δx--
		} else if Δx < 0 {
			Δx++
		}
		Δy--
	}
	return math.MinInt
}

func main() {
	start := point{x: 0, y: 0}
	target := bounds{topLeft: point{x: 179, y: -63}, bottomRight: point{x: 201, y: -109}}

	part1(start, target)
	part2(start, target)
}
