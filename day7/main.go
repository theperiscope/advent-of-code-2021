package main

import (
	"AOC/pkg/utils"
	"fmt"
)

// linearDistanceCost implements part 1 where cost is fixed at 1
func fixedDistanceCost(input []int, target int) (s int) {
	for _, n := range input {
		d := utils.AbsInt(n - target)
		s += d
	}
	return
}

// linearDistanceCost implements part 2 where cost increases from 1 by 1
func linearDistanceCost(input []int, target int) (s int) {
	for _, n := range input {
		d := utils.AbsInt(n - target)
		s += (d * (d + 1)) / 2
	}
	return
}

// align returns minimum cost to align the input elements given their initial positions
func align(input []int, travelDistanceCost func([]int, int) int) int {
	left, right := utils.MinMax(input)
	for utils.AbsInt(left-right) > 1 {
		d1 := travelDistanceCost(input, left)
		d2 := travelDistanceCost(input, right)

		if d1 < d2 {
			right = left + ((right - left) / 2)
		} else {
			left = (left + right) / 2
		}
	}

	return utils.MinInt(travelDistanceCost(input, left), travelDistanceCost(input, right))
}

func main() {
	utils.AssertArgs()
	crabs := utils.AssertInputIntCsv()

	part1 := align(crabs, fixedDistanceCost)
	fmt.Println("Part 1 Minimum Cost := ", part1)
	part2 := align(crabs, linearDistanceCost)
	fmt.Println("Part 2 Minimum Cost := ", part2)
}
