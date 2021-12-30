package main

import (
	"AOC/pkg/utils"
	"fmt"
	"math"
	"strings"
)

func part1(polymerTemplate string, rules map[string]string) {
	pairs := map[string]int{}
	for i := 0; i < len(polymerTemplate)-1; i++ {
		k := polymerTemplate[i : i+2]
		pairs[k]++
	}

	for i := 0; i < 10; i++ {
		pairs = countPairs(pairs, rules)
	}

	letters := map[string]int{}

	// first and last character are part of only one pair, but rest - two (e.g. in `AXB` the X is part of both AX and XB)
	// so we need to increase first and last letter counts by 1 before we start to normalize
	letters[string(polymerTemplate[0])]++
	letters[string(polymerTemplate[len(polymerTemplate)-1])]++
	for k, v := range pairs {
		letters[string(k[0])] += v
		letters[string(k[1])] += v
	}

	// half the values due to pair-doubling
	for k, v := range letters {
		letters[k] = v / 2
	}

	min, max := minMax(letters)
	fmt.Printf("Part 1 Answer: %d - %d = %d\n", max, min, max-min)
}

func part2(polymerTemplate string, rules map[string]string) {
	pairs := map[string]int{}
	for i := 0; i < len(polymerTemplate)-1; i++ {
		k := polymerTemplate[i : i+2]
		pairs[k]++
	}

	for i := 0; i < 40; i++ {
		pairs = countPairs(pairs, rules)
	}

	letters := map[string]int{}

	// first and last character are part of only one pair, but rest - two (e.g. in `AXB` the X is part of both AX and XB)
	// so we need to increase first and last letter counts by 1 before we start to normalize
	letters[string(polymerTemplate[0])]++
	letters[string(polymerTemplate[len(polymerTemplate)-1])]++
	for k, v := range pairs {
		letters[string(k[0])] += v
		letters[string(k[1])] += v
	}

	// half the values due to pair-doubling
	for k, v := range letters {
		letters[k] = v / 2
	}

	min, max := minMax(letters)
	fmt.Printf("Part 2 Answer: %d - %d = %d\n", max, min, max-min)
}

func minMax(letters map[string]int) (min, max int) {
	min, max = math.MaxInt, 0
	for _, v := range letters {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}

func countPairs(pairs map[string]int, rules map[string]string) map[string]int {
	pairCounts := map[string]int{}
	for k, v := range pairs {
		pairCounts[string(k[0])+rules[k]] += v
		pairCounts[rules[k]+string(k[1])] += v
	}
	return pairCounts
}

func main() {
	utils.AssertArgs()
	lines := utils.AssertInput()

	polymerTemplate := lines[0]
	rules := map[string]string{}
	for _, rule := range utils.Filter(lines, func(s string) bool { return strings.Contains(s, "->") }) {
		k, v := "", ""
		fmt.Sscanf(rule, "%s -> %s", &k, &v)
		rules[k] = v
	}

	part1(polymerTemplate, rules)
	part2(polymerTemplate, rules)
}
