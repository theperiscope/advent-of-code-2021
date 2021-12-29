package main

import (
	"AOC/pkg/utils"
	"fmt"
)

func part1(input []int) int {
	count := 0
	fmt.Println(len(input))
	for i, _ := range input {
		if i == 0 {
			fmt.Printf("%d (N/A - no previous measurement)\n", input[i])
			continue
		}

		if input[i] > input[i-1] {
			count++
			fmt.Printf("%d (increased)\n", input[i])
		} else {
			fmt.Printf("%d (decreased)\n", input[i])
		}
	}

	fmt.Println("\nHow many measurements are larger than the previous measurement?")
	fmt.Printf("> %d", count)

	return count
}

func part2(input []int) int {
	count := 0
	fmt.Println(len(input))
	for i, _ := range input {

		if i == 0 || i >= len(input)-2 {
			if i == 0 {
				fmt.Printf("N/A, %d (N/A - no previous sum)\n", input[i]+input[i+1]+input[i+2])
			}
			continue
		}

		prev := input[i-1] + input[i] + input[i+1]
		curr := input[i] + input[i+1] + input[i+2]

		if curr > prev {
			count++
			fmt.Printf("%d (increased)\n", curr)
		} else if curr == prev {
			fmt.Printf("%d (no change)\n", curr)
		} else {
			fmt.Printf("%d (decreased)\n", curr)
		}
	}

	fmt.Println("\nHow many sums are larger than the previous sum?")
	fmt.Printf("> %d", count)

	return count
}

func main() {
	utils.AssertArgs()
	input := utils.AssertInputInt()
	part1(input)
	part2(input)
}
