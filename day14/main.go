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

	p := polymerTemplate

	N := 5
	for i := 0; i < N; i++ {
		p = step(p, rules)
		fmt.Println(len(p), count(p), p)

		pairs = step2(pairs, rules)
		fmt.Println(pairs)
		if len(p) == 97 {
			pairs := map[string]int{}
			for i := 0; i < len(p)-1; i++ {
				k := p[i : i+2]
				pairs[k]++
			}
			fmt.Println(pairs)
		}
		fmt.Println("----")
	}

	counts := count(p)
	minCount, maxCount := math.MaxInt, 0
	fmt.Println(counts)
	for _, v := range counts {
		if v < minCount {
			minCount = v
		}
		if v > maxCount {
			maxCount = v
		}
	}

	fmt.Printf("Part 1 Answer: %d - %d = %d\n", maxCount, minCount, maxCount-minCount)

}

func part2(polymerTemplate string, rules map[string]string) {
	pairs := map[string]int{}
	for i := 0; i < len(polymerTemplate)-1; i++ {
		k := polymerTemplate[i : i+2]
		pairs[k]++
	}
	fmt.Println(pairs)

	N := 40
	for i := 0; i < N; i++ {
		pairs = step2(pairs, rules)
	}

	fmt.Println(pairs)

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

	fmt.Println(letters)

	minCount, maxCount := math.MaxInt, 0
	for _, v := range letters {
		if v < minCount {
			minCount = v
		}
		if v > maxCount {
			maxCount = v
		}
	}
	fmt.Printf("Part 2 Answer: %d - %d = %d\n", maxCount, minCount, maxCount-minCount)
}

func count(polymer string) map[string]int {
	m := map[string]int{}
	for i := 0; i < len(polymer); i++ {
		m[string(polymer[i])]++
	}
	return m
}

func step(polymer string, rules map[string]string) string {
	p := polymer
	insertIndexes := []int{}
	insertValues := []string{}

	for i := 0; i < len(p)-1; i++ {
		k := p[i : i+2]
		if _, ok := rules[k]; !ok {
			continue
		}

		insertIndexes = append(insertIndexes, i+1)
		insertValues = append(insertValues, rules[k])
	}

	for x := 0; x < len(insertIndexes); x++ {
		i := insertIndexes[x]
		s := insertValues[x]
		before, after := split(p, i+x /* each insert increases p so we need to modify index */)
		p = before + s + after
	}

	return p
}

func step2(pairs map[string]int, rules map[string]string) map[string]int {
	pairCounts := map[string]int{}

	for k, v := range pairs {
		s := string(k[0]) + rules[k] + string(k[1]) // length 3
		p1 := s[0:2]
		p2 := s[1:3]
		pairCounts[p1] += v
		pairCounts[p2] += v
	}

	return pairCounts
}

func split(s string, index int) (before string, after string) {
	return s[:index], s[index:]
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
